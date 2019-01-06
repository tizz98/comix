package app

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/PuerkitoBio/goquery"
)

type DilbertComicClient struct {
	BaseUrl *url.URL
}

type DilbertComic struct {
	ImageUrl string
	Title    string
}

const dilbertUrlDateFormat = "2006-01-02"

func NewDilbertComicCilent(baseUrl string) (*DilbertComicClient, error) {
	uri, err := url.Parse(baseUrl)
	return &DilbertComicClient{BaseUrl: uri}, err
}

func (c *DilbertComicClient) ComicForDate(t time.Time) (*DilbertComic, error) {
	comicUrl := c.url(fmt.Sprintf("/strip/%s", t.Format(dilbertUrlDateFormat)))

	resp, err := http.Get(comicUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse html")
	}

	s := doc.Find(".img-comic")
	if s.Length() == 0 {
		return nil, fmt.Errorf("unable to find comic in html")
	}

	imgUrl, _ := s.Eq(0).Attr("src")
	title, _ := s.Eq(0).Attr("alt")

	if !strings.HasPrefix(imgUrl, "http") {
		imgUrl = "https://" + strings.TrimLeft(imgUrl, "//")
	}

	return &DilbertComic{ImageUrl: imgUrl, Title: strings.TrimRight(title, " - Dilbert by Scott Adams")}, nil
}

func (c *DilbertComicClient) url(path string) string {
	u := *c.BaseUrl
	u.Path = strings.TrimRight(u.Path, "/")
	u.Path += path
	return u.String()
}
