package filter

import "math"

const (
	// E is 0.0000010
	E = 1.0e-6

	// Midpoint best results with 0.5
	Midpoint = 0.5
	// Factor best results with 10
	Factor = 10
)

// Clamp needs a more detailed comment ;)
// color conversion?
func Clamp(v float64) uint8 {
	return uint8(math.Min(math.Max(v, 0.0), 255.0) + 0.5)
}

// Sigmoid
// Consider replacing with https://github.com/montanaflynn/stats
func Sigmoid(a, b, x float64) float64 {
	return 1 / (1 + math.Exp(b*(a-x)))
}
