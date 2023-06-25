package filter

import "math"

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
