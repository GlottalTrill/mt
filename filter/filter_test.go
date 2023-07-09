package filter

import "testing"

func TestClamp(t *testing.T) {
	clampTests := []struct {
		input1, input2 float64
		want           uint8
	}{
		{input1: 12.34, input2: 56.78, want: 255},
		{9.9, 1.2, 12},
	}

	for _, tt := range clampTests {
		got := clamp(tt.input1 * tt.input2)
		if got != tt.want {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}

func TestSigmoid(t *testing.T) {
	var want float64 = 0

	got := sigmoid(10, 245, 3)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
