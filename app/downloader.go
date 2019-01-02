package app

import "time"

type Downloader interface {
	// Returns the comic image as bytes, title, and any error
	DownloadComic(time.Time) (*Comic, error)
}
