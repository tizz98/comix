package app

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"os/signal"
	"path"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Downloader interface {
	// Returns the comic image as bytes (png or jpg), title, and any error
	DownloadComic(time.Time) (*Comic, error)
}

type DownloaderType string

const (
	DownloaderTypeUnknown DownloaderType = "unknown"
	DownloaderTypeXkcd    DownloaderType = "xkcd"
)

type DownloaderContext struct {
	Type         DownloaderType
	LastDownload *time.Time

	outputFileDirectory string

	m sync.Mutex
}

func (ctx *DownloaderContext) Run(t time.Time) error {
	ctx.m.Lock()
	defer ctx.m.Unlock()

	if !ctx.shouldDownloadNew(t) {
		logrus.Debug("same date, not downloading new comic")
		return nil
	}

	ctx.LastDownload = &t

	downloader := ctx.downloader()
	if comic, err := downloader.DownloadComic(t); err != nil || comic == nil {
		return err
	} else {
		f, err := os.Create(ctx.filePath(t))
		if err != nil {
			return err
		}
		defer f.Close()

		mimeType := http.DetectContentType(comic.ImageData)
		switch mimeType {
		case "image/png":
		case "image/jpeg":
			img, err := jpeg.Decode(bytes.NewReader(comic.ImageData))
			if err != nil {
				return err
			}

			comic.ImageData = []byte{}
			if err := png.Encode(bytes.NewBuffer(comic.ImageData), img); err != nil {
				return err
			}
		default:
			return fmt.Errorf("invalid image type: %#v", mimeType)
		}

		// todo: save title somewhere?
		if _, err := f.Write(comic.ImageData); err != nil {
			return err
		}

		logrus.Debugf("Saved %#v to %#v", comic.Title, f.Name())
	}

	return nil
}

func (ctx *DownloaderContext) downloader() Downloader {
	if ctx.Type == DownloaderTypeXkcd {
		return &XkcdDownloader{}
	}
	return nil
}

func (ctx *DownloaderContext) shouldDownloadNew(t time.Time) bool {
	return ctx.LastDownload == nil || !isSameDate(t, *ctx.LastDownload)
}

func (ctx *DownloaderContext) filePath(t time.Time) string {
	return path.Join(ctx.outputFileDirectory, fmt.Sprintf("%s_comic_%s.png", ctx.Type, t.Format(time.RFC3339Nano)))
}

func isSameDate(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

func RunDownloader(downloaderType DownloaderType, outputPath string) error {
	if downloaderType == DownloaderTypeUnknown {
		return fmt.Errorf("invalid download type")
	}

	signalChan := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(signalChan, os.Interrupt)

	ticker := time.NewTicker(time.Minute)
	ctx := DownloaderContext{Type: downloaderType, outputFileDirectory: outputPath}

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
