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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &Comic{ImageData: body, Title: comic.Title}, nil
}

func isSameDate(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}
