package plot

import (
	"testing"
	"fmt"
)

func TestDraw(t *testing.T) {
	a := []float64{3, 7, 7}
	b := []float64{2, 6, 8}
	c := []float64{1, 2, 3}
	e := []float64{10, 6, 3}
	f := []float64{1, 16, 23}

	Draw(a, b, c,e,f)
}

func TestCalcOffset(t *testing.T) {
	a := CalcOffset(4)
	fmt.Println(a)
}