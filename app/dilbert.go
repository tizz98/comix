package app

import (
	"io/ioutil"
	"net/http"
	"time"
)

type DilbertDownloader struct {
}

func (d *DilbertDownloader) DownloadComic(t time.Time) (*Comic, error) {
	client, err := NewDilbertComicCilent("https://dilbert.com")
	if err != nil {
		return nil, err
	}

	comic, err := client.ComicForDate(t)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(comic.ImageUrl)
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
