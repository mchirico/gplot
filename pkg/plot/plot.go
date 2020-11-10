package plot

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"image/color"
)

func Convert(v []float64) plotter.Values {
	c := make(plotter.Values, len(v))
	c = v
	return c
}

func Bar(group plotter.Values, w vg.Length, i int) *plotter.BarChart {
	bar, err := plotter.NewBarChart(group, vg.Points(10))
	if err != nil {
		panic(err)
	}
	bar.LineStyle.Width = vg.Length(0)
	bar.Color = plotutil.Color(i)
	bar.Offset = w
	return bar
}

func Draw(v ...[]float64) {

	group := []plotter.Values{}
	for i,_ := range v {
		group = append(group,v[i])
	}


	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.BackgroundColor = color.RGBA{R: 12, B: 28, G: 50, A: 55}

	p.Title.Text = "Bar chart"
	p.Y.Label.Text = "Heights"

	offset := []vg.Length{-20,-10,0,10}
	bars := []plot.Plotter{}

	for i,_ := range group {
		bars = append(bars,Bar(group[i],offset[i],i ))

	}

	p.Add(bars...)
	names := []string{}
	for i,_ := range bars {
		names = append(names,fmt.Sprintf("Group %d",i) )
		p.Legend.Add(names[i], bars[i].(plot.Thumbnailer))
	}


	p.Legend.Top = true
	p.NominalX(names...)

	if err := p.Save(5*vg.Inch, 3*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}
	if err := p.Save(5*vg.Inch, 3*vg.Inch, "barchart.svg"); err != nil {
		panic(err)
	}
}
