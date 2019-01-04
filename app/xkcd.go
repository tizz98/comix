package app

import (
	"io/ioutil"
	"net/http"
	"time"

	xkcd "github.com/nishanths/go-xkcd"
)

type XkcdDownloader struct {
}

func (x *XkcdDownloader) DownloadComic(t time.Time) (*Comic, error) {
	var comic xkcd.Comic

	client := xkcd.NewClient()
	comic, err := client.Latest()
	if err != nil {
		return nil, err
	}

	if !isSameDate(t, comic.PublishDate) {
		comic, err = client.Random()
		if err != nil {
			return nil, err
		}
	}

	resp, err := http.Get(comic.ImageURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Comic{ImageData: body, Title: comic.Title}, nil
}
