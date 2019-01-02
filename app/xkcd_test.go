package app

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestXkcdDownloader_DownloadComic(t *testing.T) {
	d := time.Date(2018, 12, 31, 0, 0, 0, 0, time.Local)
	dl := &XkcdDownloader{}

	comic, err := dl.DownloadComic(d)
	require.NoError(t, err)

	assert.NotEmpty(t, comic.Title)
	assert.NotNil(t, comic.ImageData)
}
