package main

import (
	"bytes"
	"github.com/disintegration/imaging"
	xfilter "github.com/mutschler/mt/filter"
	"image"
)

// ImageStripFilter wraps AddStripsToImage because Asset() is declared in package main
// and will take longer to refactor out.
// DEPRECATION: AddStripsToImage will replace this
func ImageStripFilter(img image.Image) image.Image {
	l, _ := Asset("strip_left.jpg")
	lr := bytes.NewReader(l)
	strip, _ := imaging.Decode(lr)

	r, _ := Asset("strip_right.jpg")
	rr := bytes.NewReader(r)
	stripr, _ := imaging.Decode(rr)

	modifiedImg := xfilter.AddStripsToImage(img, strip, stripr)

	return modifiedImg
}

func CrossProcessingFilter(img image.Image) *image.NRGBA {
	return xfilter.CrossProcessing(img)
}
