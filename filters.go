package main

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/mutschler/mt/filter"
	"image"
	"image/color"
	"math"
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

	modifiedImg := filter.AddStripsToImage(img, strip, stripr)

	return modifiedImg
}

// CrossProcessingFilter wraps the sigmoid function to simulate image cross processing.
// Best results with midpoint: 0.5 and factor 10
func CrossProcessingFilter(img image.Image) *image.NRGBA {
	// TODO:  move these to a colours package?
	red := make([]uint8, 256)
	green := make([]uint8, 256)
	blue := make([]uint8, 256)
	a := math.Min(math.Max(filter.Midpoint, 0.0), 1.0)
	b := math.Abs(filter.Factor)
	sig0 := filter.Sigmoid(a, b, 0)
	sig1 := filter.Sigmoid(a, b, 1)

	for i := 0; i < 256; i++ {
		x := float64(i) / 255.0
		sigX := filter.Sigmoid(a, b, x)
		f := (sigX - sig0) / (sig1 - sig0)
		red[i] = filter.Clamp(f * 255.0)
	}

	for i := 0; i < 256; i++ {
		x := float64(i) / 255.0
		sigX := filter.Sigmoid(a, b, x)
		f := (sigX - sig0) / (sig1 - sig0)
		green[i] = filter.Clamp(f * 255.0)
	}

	for i := 0; i < 256; i++ {
		x := float64(i) / 255.0
		arg := math.Min(math.Max((sig1-sig0)*x+sig0, filter.E), 1.0-filter.E)
		f := a - math.Log(1.0/arg-1.0)/b
		blue[i] = filter.Clamp(f * 255.0)
	}

	fn := func(c color.NRGBA) color.NRGBA {
		return color.NRGBA{R: red[c.R], G: green[c.G], B: blue[c.B], A: c.A}
	}

	return imaging.AdjustFunc(img, fn)
}
