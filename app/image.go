package app

import (
	"image"
	"time"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/inconsolata"
)

const textHeight = 25

// Creates a new image with the date `t` at the top.
// todo: handle different aspect ratios
func makeImage(comicImg image.Image, t time.Time) image.Image {
	comicSize := comicImg.Bounds().Size()

	width := comicSize.X
	height := comicSize.Y + textHeight

	dc := gg.NewContext(comicSize.X, textHeight)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetFontFace(inconsolata.Regular8x16)
	dc.DrawStringAnchored(t.Format("Mon, 02 Jan 2006"), float64(comicSize.X)/2, textHeight/2, 0.5, 0.5)
	textImg := dc.Image()

	dc = gg.NewContext(width, height)
	dc.DrawImage(textImg, 0, 0)
	dc.DrawImage(comicImg, 0, textHeight)

	return dc.Image()
}
