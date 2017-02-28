package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkewStr(t *testing.T) {
	assert.Equal(t,
		[]int{0, -1, -1, -1, 0, 1, 2, 1, 1, 1, 0, 1, 2, 1, 0, 0, 0, 0, -1, 0, -1, -2},
		SkewStr("CATGGGCATCGGCCATACGCC"))

	assert.Equal(t,
		[]int{0, 1, 1, 2, 1, 0, 0, -1, -2, -1, -2, -1, -1, -1, -1},
		SkewStr("GAGCCACCGCGATA"))

	skewData := SkewStr("TAAAGACTGCCGAGAGGCCAACACGAGTGCTAGAACGAGGGGCGTAAACGCGGGTCCGAT")
	positions, value := Minimum(skewData)
	assert.Equal(t, []int{11, 24}, positions)
	assert.Equal(t, -1, value)

	data, err := ioutil.ReadFile("dataset_7_6.txt")
	if err != nil {
		panic(err)
	}
	skewData = SkewStr(string(data))
	positions, value = Minimum(skewData)
	assert.Equal(t, []int{47170, 47171}, positions)
	assert.Equal(t, -313, value)
}

func TestSkewPlot(t *testing.T) {
	data, err := ioutil.ReadFile("E_coli.txt")
	if err != nil {
		panic(err)
	}
	dna := string(data)

	_, err = SkewPlot("E. Coli", "ecoli_plot.png", dna)
	if err != nil {
		panic(err)
	}
}
