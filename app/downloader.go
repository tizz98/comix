package app

import (
	"bytes"
	"context"
	"fmt"
	"image/png"
	"net/http"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"
	"time"

	"github.com/inconshreveable/go-update"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/tizz98/comix/cnc"
)

type ComicDownloader interface {
	// Returns the comic image as bytes (png or jpg), title, and any error
	DownloadComic(time.Time) (*Comic, error)
}

type DownloaderType string

const (
	DownloaderTypeUnknown DownloaderType = "unknown"
	DownloaderTypeXkcd    DownloaderType = "xkcd"
	DownloaderTypeDilbert DownloaderType = "dilbert"
)

type DownloaderContext struct {
	Type         DownloaderType
	LastDownload *time.Time

	outputFileDirectory string
	cncAddr             string
	cncClientId         string

	m        sync.Mutex
	stopChan chan os.Signal
}

func (ctx *DownloaderContext) Run(t time.Time) error {
	// todo: refactor this monster function
	ctx.m.Lock()
	defer ctx.m.Unlock()

	if !ctx.shouldDownloadNew(t) {
		logrus.Debug("same date, not downloading new comic")
		ctx.updateBinary()
		return nil
	}

	ctx.LastDownload = &t

	downloader := ctx.downloader()
	if comic, err := downloader.DownloadComic(t); err != nil || comic == nil {
		logrus.WithError(err).Error("unable to download comic")
		return err
	} else {
		f, err := os.Create(ctx.filePath(t))
		if err != nil {
			return err
		}
		defer f.Close()

		comic.ImageData, err = ToPng(comic.ImageData)
		if err != nil {
			logrus.WithError(err).Error("unable to convert to png")
			return err
		}

		img, err := png.Decode(bytes.NewReader(comic.ImageData))
		if err != nil {
			logrus.WithError(err).Error("unable to decode png")
			return err
		}

		composedImg := makeImage(img, t)
		if err := png.Encode(f, composedImg); err != nil {
			logrus.WithError(err).Error("unable to encode png")
			return err
		}

		logrus.Infof("Saved %#v to %#v", comic.Title, f.Name())
	}

	// todo: enable screen update
	//ctx.updateScreen()

	// if an update happened, stop the program.
	// a supervisor should restart it.
	if didUpdate := ctx.updateBinary(); didUpdate {
		ctx.stopChan <- syscall.SIGINT
	}

	return nil
}

func (ctx *DownloaderContext) updateBinary() bool {
	if ctx.cncAddr == "" || ctx.cncClientId == "" {
		logrus.Warn("cnc address or client id not set, not performing automatic updates")
		return false
	}

	conn, err := grpc.Dial(ctx.cncAddr, grpc.WithInsecure())
	if err != nil {
		logrus.WithError(err).Error("unable to connect to grpc server")
		return false
	}
	defer conn.Close()

	client := cnc.NewCnCClient(conn)
	resp, err := client.Ping(context.Background(), &cnc.PingMsg{
		ClientId: ctx.cncClientId,
		Ok:       true,
	})
	if err != nil {
		logrus.WithError(err).Error("unable to ping cnc server")
		return false
	}

	errorMsg := ""

	defer func() {
		logrus.Debug("sending status update")
		if _, err := client.UpdateStatus(context.Background(), &cnc.UpdateMsg{
			ClientId:       ctx.cncClientId,
			UpdateMessage:  errorMsg,
			UpdateComplete: errorMsg == "",
		}); err != nil {
			logrus.WithError(err).Error("unable to post status update")
		} else {
			logrus.Debug("finished sending status update")
		}
	}()

	if resp.HasUpdate {
		logrus.Debug("downloading updated binary")

		httpResp, err := http.Get(resp.GetUrl())
		if err != nil {
			logrus.WithError(err).Error("unable to request binary url")
			return false
		}
		defer httpResp.Body.Close()

		logrus.Debug("applying update")
		err = update.Apply(httpResp.Body, update.Options{})
		if err != nil {
			logrus.WithError(err).Debug("error applying update, trying to rollback")

			if rerr := update.RollbackError(err); rerr != nil {
				logrus.WithError(rerr).Error("failed to rollback from bad update")
			} else {
				logrus.Debug("rollback succeeded")
			}

			return false
		}

		logrus.Debug("finished applying update")
		return true
	} else {
		logrus.Debug("no binary update")
	}

	return false
}

// if any error happens in this method, it is logged and not thrown so that the program does not
// exit and the current image being displayed will still be displayed.
func (ctx *DownloaderContext) updateScreen() {
	p := ctx.filePath(time.Now())
	f, err := os.Open(p)

	if err != nil {
		logrus.WithError(err).Error("unable to open file")
		return
	}

	img, err := png.Decode(f)
	if err != nil {
		logrus.WithError(err).Error("unable to decode png")
		return
	}

	d := NewEpdDisplayer()
	if err := d.WriteImage(img); err != nil {
		logrus.WithError(err).Error("unable to write image to epd")
		return
	}
}

func (ctx *DownloaderContext) downloader() ComicDownloader {
	switch ctx.Type {
	case DownloaderTypeXkcd:
		return &XkcdDownloader{}
	case DownloaderTypeDilbert:
		return &DilbertDownloader{}
	}
	return nil
}

func (ctx *DownloaderContext) shouldDownloadNew(t time.Time) bool {
	return ctx.LastDownload == nil || !isSameDate(t, *ctx.LastDownload)
}

func (ctx *DownloaderContext) filePath(t time.Time) string {
	return path.Join(ctx.outputFileDirectory, fmt.Sprintf("%s_comic_%s.png", ctx.Type, t.Format("20060102")))
}

func isSameDate(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

type Option struct {
	TickDuration *time.Duration
	CnCAddress   string
	ClientId     string
}

func NewDuration(d time.Duration) *time.Duration {
	return &d
}

func RunDownloader(downloaderType DownloaderType, outputPath string, options *Option) error {
	if downloaderType == DownloaderTypeUnknown {
		return fmt.Errorf("invalid download type")
	}

	signalChan := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)

	ticker := time.NewTicker(time.Minute)
	ctx := DownloaderContext{Type: downloaderType, outputFileDirectory: outputPath, stopChan: signalChan}

	if options != nil {
		ticker = time.NewTicker(*options.TickDuration)

		ctx.cncAddr = options.CnCAddress
		ctx.cncClientId = options.ClientId
	}

	go func() {
		for {
			select {
			case <-signalChan:
				logrus.Info("received interrupt, shutting down...")
				ticker.Stop()
				close(done)
			case <-ticker.C:
				logrus.Debug("received tick")
				go ctx.Run(time.Now())
			}
		}
	}()

	<-done
	return nil
}
