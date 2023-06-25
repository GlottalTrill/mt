package main

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/mutschler/mt/filter"
	"image"
	"image/color"
	"math"
)

// ImageStrip look
func ImageStripFilter(img image.Image) image.Image {
	l, _ := Asset("strip_left.jpg")
	lr := bytes.NewReader(l)
	strip, _ := imaging.Decode(lr)
	strip = imaging.Resize(strip, 0, img.Bounds().Dy(), imaging.Lanczos)

	dst := imaging.New((2*strip.Bounds().Dx())+img.Bounds().Dx(), img.Bounds().Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, img, image.Pt(strip.Bounds().Dx(), 0))
	dst = imaging.Paste(dst, strip, image.Pt(0, 0))

	r, _ := Asset("strip_right.jpg")
	rr := bytes.NewReader(r)
	stripr, _ := imaging.Decode(rr)
	stripr = imaging.Resize(stripr, 0, img.Bounds().Dy(), imaging.Lanczos)
	dst = imaging.Paste(dst, stripr, image.Pt(dst.Bounds().Dx()-stripr.Bounds().Dx(), 0))
	return dst
}

// sigmoid function to simulate image cross processing, best results with midpoint: 0.5 and factor 10
func CrossProcessingFilter(img image.Image, midpoint, factor float64) *image.NRGBA {

	// TODO:  move these to a colours package?
	red := make([]uint8, 256)
	green := make([]uint8, 256)
	blue := make([]uint8, 256)
	a := math.Min(math.Max(midpoint, 0.0), 1.0)
	b := math.Abs(factor)
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
		return color.NRGBA{red[c.R], green[c.G], blue[c.B], c.A}
	}

	return imaging.AdjustFunc(img, fn)
}
