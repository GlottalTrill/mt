package filter

import (
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"math"
)

const (
	// E is 0.0000010
	E = 1.0e-6

	// Midpoint best results with 0.5
	Midpoint = 0.5
	// Factor best results with 10
	Factor = 10
)

// AddStripsToImage adds "filmstrips" to the left and right sides of a passed image
func AddStripsToImage(img, leftStrip, rightStrip image.Image) image.Image {
	// resize the "filmstrip" to match the height of the passed image
	leftStrip = imaging.Resize(leftStrip, 0, img.Bounds().Dy(), imaging.Lanczos)
	rightStrip = imaging.Resize(rightStrip, 0, img.Bounds().Dy(), imaging.Lanczos)

	dst := imaging.New((2*leftStrip.Bounds().Dx())+img.Bounds().Dx(), img.Bounds().Dy(), color.NRGBA{})
	dst = imaging.Paste(dst, img, image.Pt(leftStrip.Bounds().Dx(), 0))
	dst = imaging.Paste(dst, leftStrip, image.Pt(0, 0))
	dst = imaging.Paste(dst, rightStrip, image.Pt(dst.Bounds().Dx()-rightStrip.Bounds().Dx(), 0))
	return dst
}

// CrossProcessing wraps the sigmoid function to simulate image cross processing.
// Best results with midpoint: 0.5 and factor 10
func CrossProcessing(img image.Image) *image.NRGBA {
	// TODO:  move these to a colours package?
	red := make([]uint8, 256)
	green := make([]uint8, 256)
	blue := make([]uint8, 256)
	a := math.Min(math.Max(Midpoint, 0.0), 1.0)
	b := math.Abs(Factor)
	sig0 := sigmoid(a, b, 0)
	sig1 := sigmoid(a, b, 1)

	for i := 0; i < 256; i++ {
		x := float64(i) / 255.0
		sigX := sigmoid(a, b, x)
		f := (sigX - sig0) / (sig1 - sig0)
		red[i] = clamp(f * 255.0)
	}

	for i := 0; i < 256; i++ {
		x := float64(i) / 255.0
		sigX := sigmoid(a, b, x)
		f := (sigX - sig0) / (sig1 - sig0)
		green[i] = clamp(f * 255.0)
	}

	for i := 0; i < 256; i++ {
		x := float64(i) / 255.0
		arg := math.Min(math.Max((sig1-sig0)*x+sig0, E), 1.0-E)
		f := a - math.Log(1.0/arg-1.0)/b
		blue[i] = clamp(f * 255.0)
	}

	fn := func(c color.NRGBA) color.NRGBA {
		return color.NRGBA{R: red[c.R], G: green[c.G], B: blue[c.B], A: c.A}
	}

	return imaging.AdjustFunc(img, fn)
}

// clamp needs a more detailed comment ;)
// color conversion?
func clamp(v float64) uint8 {
	return uint8(math.Min(math.Max(v, 0.0), 255.0) + 0.5)
}

// sigmoid
// Consider replacing with https://github.com/montanaflynn/stats
func sigmoid(a, b, x float64) float64 {
	return 1 / (1 + math.Exp(b*(a-x)))
}
