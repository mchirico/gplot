package plot

import "testing"

func TestDraw(t *testing.T) {
	a := []float64{3, 5, 7}
	b := []float64{2, 4, 8}
	c := []float64{1, 7, 3}

	Draw(a, b, c)
}
