package plot

import "testing"

func TestDraw(t *testing.T) {
	a := []float64{3, 7, 7}
	b := []float64{2, 6, 8}
	c := []float64{1, 2, 3}

	Draw(a, b, c)
}
