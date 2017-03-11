package main

import (
	"log"
	"time"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

// DNA is an uppercase string of ACGT
func SkewStr(dna string) []int {
	return Skew(NormalizeDNA(dna))
}

// DNA is a byte slice of normalized DNA to values 0-4
func Skew(dna []byte) []int {
	dnaLen := len(dna)
	skew := make([]int, dnaLen+1, dnaLen+1)

	skew[0] = 0
	for i, bp := range dna {
		skew[i+1] = skew[i] + skewValue(byte(bp))
	}
	return skew
}

func MinSkew(dna []byte) (positions []int, value int) {
	return Minimum(Skew(dna))
}

func SkewPlot(title, filename, dna string) ([]int, error) {
	start := time.Now()
	skewData := SkewStr(dna)
	log.Print("skewData done in ", time.Since(start).String())

	start = time.Now()
	plotPoints := createPlotPoints(skewData)
	p, err := plot.New()
	if err != nil {
		return nil, err
	}

	p.Title.Text = title + " Skew(G-C) plot"
	p.X.Label.Text = "bp"
	p.Y.Label.Text = "G-C (Y)"

	err = plotutil.AddLinePoints(p,
		title, plotPoints)
	if err != nil {
		return nil, err
	}
	log.Print("plotting done in ", time.Since(start).String())

	// Save the plot to a PNG file.
	start = time.Now()
	if err := p.Save(30*vg.Centimeter, 20*vg.Centimeter, filename); err != nil {
		return nil, err
	}
	log.Print("saving plot done in ", time.Since(start).String())

	return skewData, nil
}

func createPlotPoints(data []int) plotter.XYs {
	pts := make(plotter.XYs, len(data))
	for i := range pts {
		pts[i].X = float64(i)
		pts[i].Y = float64(data[i])
	}
	return pts
}

var skewValues = []int{
	0: 0,  // A
	1: -1, // C
	2: 1,  // G
	3: 0,  // T
}

func skewValue(bp byte) int {
	return int(skewValues[bp])
}
