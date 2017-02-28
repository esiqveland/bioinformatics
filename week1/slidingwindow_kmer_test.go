package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingWindowFrequentWordsFaster(t *testing.T) {
	data := "CGGACTCGACAGATGTGAAGAACGACAATGTGAAGACTCGACACGACAGAGTGAAGAGAAGAGGAAACATTGTAA"
	dna := strings.ToUpper(data)

	results := MovingWindowFrequentWordsFaster(dna, 5, 50, 4)

	fmt.Printf("MovingWindowFrequentWords(dna, 5, 50, 4): %+v\n", results)
	assert.Contains(t, results, "CGACA")
	assert.Contains(t, results, "GAAGA")
}

func TestMovingWindowFrequentWordsFaster_big(t *testing.T) {
	data, err := ioutil.ReadFile("test_pat_count.txt")
	if err != nil {
		t.Errorf("error reading test data: %v", err)
	}

	dna := strings.ToUpper(string(data))

	results := MovingWindowFrequentWordsFaster(dna, 11, 566, 18)

	fmt.Printf("MovingWindowFrequentWords(dna, 11, 566, 18): %+v\n", results)
	assert.Contains(t, results, "AAACCAGGTGG")
}

// MovingWindowFrequentWordsFaster
func TestMovingWindowFrequentWordsFaster_bigger(t *testing.T) {
	data, err := ioutil.ReadFile("dataset_4_5.txt")
	if err != nil {
		t.Errorf("error reading test data: %v", err)
	}

	dna := strings.ToUpper(string(data))

	results := MovingWindowFrequentWordsFaster(dna, 10, 481, 18)

	fmt.Printf("MovingWindowFrequentWordsFaster(dna, 10, 481, 18): %+v\n", results)
	assert.Contains(t, results, "TGGCCCTTAC")
	assert.Contains(t, results, "TAGTAAGCCT")
	assert.Contains(t, results, "CGAAAAACCG")
	assert.Contains(t, results, "CCGGCCTTAA")
	assert.Contains(t, results, "CGGTGCCCCC")
	assert.Len(t, results, 5)
}

func TestMovingWindowFrequentWords_e_coli(t *testing.T) {
	data, err := ioutil.ReadFile("E_coli.txt")
	if err != nil {
		t.Errorf("error reading test data: %v", err)
	}

	dna := strings.ToUpper(string(data))

	results := MovingWindowFrequentWordsFaster(dna, 9, 500, 3)

	fmt.Printf("E_coli results: %v\n", len(results))
	assert.Equal(t, 1904, len(results))
	//fmt.Printf("MovingWindowFrequentWords(dna, 11, 566, 18): %+v\n", results)
}

func TestFasterFrequentWords(t *testing.T) {
	wanted := []FreqWordResult{
		{Pattern: "CATG", Count: 3},
		{Pattern: "GCAT", Count: 3},
	}
	results := FasterFrequentWordsStr("ACGTTGCATGTCGCATGATGCATGAGAGCT", 4)

	assert.Equal(t, len(wanted), len(results), "Wanted %v results but got %v", len(wanted), len(results))

	for _, w := range wanted {
		assert.Contains(t, results, w.Pattern)
		assert.Equal(t, w.Count, results[w.Pattern].Count)
	}
}

func TestFrequencyCounting(t *testing.T) {
	expected := []int{2, 1, 0, 0, 0, 0, 2, 2, 1, 2, 1, 0, 0, 1, 1, 0}

	idx := NewIndexStr("ACGCGGCTCTGAAA", 2)

	assert.Equal(t, expected, idx.Frequencies())
}

func TestFrequencyCounting_big(t *testing.T) {
	data, _ := ioutil.ReadFile("dataset_2994_5.txt")
	dna := string(data)

	idx := NewIndexStr(dna, 5)
	freqs := idx.Frequencies()

	res := ""
	for _, val := range freqs {
		res += fmt.Sprintf("%v ", val)
	}
	res = res[:len(res)-1]
	assert.Equal(t, expected_2994_5, res)
}

const expected_2994_5 = "0 3 0 0 1 1 1 1 0 1 1 0 1 0 1 1 0 1 2 0 1 0 0 1 0 1 1 0 0 1 1 0 1 0 0 1 0 1 0 1 1 0 2 2 1 2 0 2 3 3 1 1 1 0 0 0 0 0 1 1 2 2 0 0 0 0 2 1 2 1 2 0 2 3 0 1 1 0 0 0 2 0 0 0 0 1 0 0 1 1 0 1 0 1 1 0 1 0 1 2 0 0 0 2 1 1 1 0 0 3 2 0 0 0 0 0 0 1 0 0 1 1 1 0 1 0 0 1 1 0 0 4 1 0 1 0 0 0 1 0 1 1 1 1 0 0 0 0 0 1 2 0 1 1 1 0 0 3 0 0 0 0 1 0 1 0 1 0 0 1 1 2 0 0 1 1 1 0 0 1 2 0 1 0 1 0 0 1 1 2 1 0 1 1 3 1 2 0 1 0 1 0 1 0 0 1 1 0 0 0 2 0 0 0 1 1 0 1 0 1 1 2 0 1 0 1 0 0 1 1 1 1 0 0 1 1 1 3 2 0 2 0 0 0 0 1 1 0 0 0 0 1 0 0 0 0 0 0 0 1 1 0 1 1 1 1 0 1 1 1 0 1 1 0 3 1 1 1 2 0 0 0 0 2 0 0 1 2 2 0 0 2 0 1 3 0 0 0 0 0 1 0 0 1 1 0 0 0 0 0 1 0 0 1 1 0 0 0 0 0 1 2 0 1 0 0 0 1 0 0 0 0 0 0 0 0 1 0 0 0 0 0 0 1 0 0 0 0 1 2 0 1 0 1 1 1 1 0 2 2 0 0 0 1 1 0 0 0 0 3 0 0 1 1 1 2 1 0 1 0 2 2 1 0 1 0 1 0 1 0 1 0 3 0 1 0 0 4 3 0 0 0 0 1 0 0 1 0 0 1 2 1 1 2 1 2 2 0 0 1 1 0 0 2 0 1 2 0 2 0 1 1 0 1 1 0 0 0 2 1 1 1 0 1 1 0 2 2 0 0 0 1 0 1 2 0 1 1 0 0 1 0 0 0 1 1 0 1 1 1 1 3 1 0 3 0 0 2 4 0 1 1 1 1 1 1 1 1 1 3 0 1 1 2 1 0 2 1 2 0 1 1 0 0 0 1 0 2 1 1 1 0 1 0 2 1 1 0 0 0 0 0 1 1 3 0 1 1 1 2 0 0 0 0 1 0 1 0 0 3 0 0 0 0 1 1 1 1 0 0 0 0 0 1 2 0 0 1 0 0 1 0 1 1 1 2 1 3 1 2 0 1 0 0 1 0 0 1 1 1 1 2 0 1 0 1 0 1 0 1 1 0 0 0 0 0 1 0 0 0 1 2 0 0 0 1 0 1 0 1 1 0 0 0 2 1 0 2 1 3 0 0 1 3 0 0 0 1 2 2 1 1 1 1 0 2 1 0 2 0 1 1 0 1 0 0 0 0 1 0 0 0 0 1 0 0 1 2 1 1 0 0 0 0 0 1 1 2 0 0 2 1 0 0 0 0 3 0 1 0 0 1 1 2 1 1 0 3 1 1 0 1 0 2 0 1 1 1 1 1 1 3 1 1 1 0 2 1 1 0 2 0 0 2 0 0 0 0 3 0 0 2 0 0 1 2 1 1 2 0 0 2 1 1 3 1 0 1 1 1 2 0 1 0 0 0 0 2 1 2 0 1 0 2 0 1 0 1 3 2 2 0 2 0 1 1 0 1 2 1 0 1 0 1 0 0 1 0 3 3 3 0 0 1 1 2 1 0 0 0 0 1 3 1 2 0 0 0 1 0 1 1 0 0 0 1 0 2 0 1 0 0 0 0 2 1 1 0 0 0 0 0 0 1 0 1 0 4 0 0 0 0 0 1 0 0 2 1 0 1 2 0 0 0 0 0 1 0 1 1 0 0 0 0 0 3 1 2 1 0 2 1 1 3 1 0 1 4 0 1 0 1 2 0 1 0 1 0 0 2 1 0 2 0 0 2 2 2 1 2 3 2 1 0 0 2 1 0 1 0 1 1 2 0 0 0 1 1 2 1 0 0 2 2 1 0 0 0 0 2 1 0 2 1 0 1 1 0 1 0 0 0 0 0 2 1 0 1 0 1 0 2 2 1 2 1 1 1 0 3 1 4 0 1 1 1 0 1 0 0 2 0 2 1 1 0 1 1 0 0 0 3 0 0 1 0 0 1 0 0 0 0 1 2 3 1 0 0 1 1 1 1 1 1 0 1 1 0 1 0 0 0 1 1 2 1 0 0 1 0 1 0 0 0 0 2 1 1 0 0 0 0 1 1"
