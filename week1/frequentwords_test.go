package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequentWordsWithMismatches(t *testing.T) {

	seqs := FrequentWordsWithMismatches(NormalizeDNA("ACGTTGCATGTCGCATGATGCATGAGAGCT"), 4, 1)

	dnaStrains := sequences(seqs).DeNormalize()

	assert.Equal(t, []string{"ATGC", "ATGT", "GATG"}, dnaStrains)
}

func TestFrequentWordsWithMismatchesRevComplement(t *testing.T) {

	seqs := FrequentWordsWithMismatchesRevComplement(NormalizeDNA("ACGTTGCATGTCGCATGATGCATGAGAGCT"), 4, 1)

	dnaStrains := sequences(seqs).DeNormalize()

	assert.Equal(t, []string{"ATGT", "ACAT"}, dnaStrains)
}

const dataset_9_7 = "CACAGCGGTAAAGCGGGCGGGCGGTAAACACACGGCACACACAGCGGCACACGGGCGGTAAAGCGGTAAAGCGGCACACACACACATAAACGGTAAACGGGTTGCGGTAAAGTTCGGCACAGCGGTAAAGTTCACACGGTAAAGTTTAAATAAAGCGGCACACGGCACACACACGGGCGGCACAGCGGCGGGCGGCACACGGTAAACGGCGGCACATAAATAAACGGCGGGTTGCGGCGGCACAGTTCACATAAATAAATAAAGCGGGTTCGGGTTTAAAGCGGGTTCACATAAACACACACAGTTCACAGCGGCGGCGGCACATAAAGCGGTAAACACACGGTAAAGCGGGTTGCGGGTTGTTCACATAAAGTTCACAGTTCACACGGTAAAGCGG"

func TestFrequentWordsWithMismatchesLong(t *testing.T) {

	seqs := FrequentWordsWithMismatches(NormalizeDNA(dataset_9_7), 7, 3)

	dnaStrains := sequences(seqs).DeNormalize()

	assert.Equal(t, []string{"AAAAAGG"}, dnaStrains)
	fmt.Printf("%v", dnaStrains)
}

func TestFrequentWordsWithMismatchesLonger(t *testing.T) {
	fasta, err := ReadFasta("Salmonella_enterica.txt")
	assert.NoError(t, err)

	seqs := FrequentWordsWithMismatches(NormalizeDNA(fasta.Genome()), 7, 3)

	dnaStrains := sequences(seqs).DeNormalize()

	assert.Equal(t, []string{"AAAAAGG"}, dnaStrains)
	fmt.Printf("%v", dnaStrains)
}

func TestFrequentWordsWithMismatchesLonger16S3(t *testing.T) {
	fasta, err := ReadFasta("fasta/16_S3.fasta")
	assert.NoError(t, err)

	seqs := FrequentWordsWithMismatches(NormalizeDNA(fasta.Genome()), 12, 1)

	dnaStrains := sequences(seqs).DeNormalize()

	assert.Equal(t, []string{"AAAAAGG"}, dnaStrains)
	fmt.Printf("%v", dnaStrains)
}
