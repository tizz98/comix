package app

import (
	"image"

	epdfuse "github.com/wmarbut/go-epdfuse"
)

type EpdDisplayer struct {
	fuse epdfuse.EpdFuse
}

func NewEpdDisplayer() *EpdDisplayer {
	return &EpdDisplayer{
		fuse: epdfuse.NewEpdFuse(),
	}
}

func (d *EpdDisplayer) WriteText(text string) error {
	return d.fuse.WriteText(text)
}

func (d *EpdDisplayer) WriteImage(img image.Image) error {
	return d.fuse.WriteImage(img)
}

func (d *EpdDisplayer) Update() error {
	return d.fuse.PartialUpdate()
}
