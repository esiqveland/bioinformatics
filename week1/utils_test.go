package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFasta(t *testing.T) {
	fasta, err := ReadFasta("Salmonella_enterica.txt")
	assert.NoError(t, err)

	fmt.Println(fasta.Raw())
	//fmt.Println(fasta.Genome())
}

func TestNormalizeDNA(t *testing.T) {
	assert.Equal(t, []byte{0}, NormalizeDNA("A"))
	assert.Equal(t, []byte{1}, NormalizeDNA("C"))
	assert.Equal(t, []byte{2}, NormalizeDNA("G"))
	assert.Equal(t, []byte{3}, NormalizeDNA("T"))
}

func TestRevComplement(t *testing.T) {
	assert.Equal(t, "T", DeNormalizeDNA(RevComplement(NormalizeDNA("A"))))
	assert.Equal(t, "A", DeNormalizeDNA(RevComplement(NormalizeDNA("T"))))
	assert.Equal(t, "C", DeNormalizeDNA(RevComplement(NormalizeDNA("G"))))
	assert.Equal(t, "G", DeNormalizeDNA(RevComplement(NormalizeDNA("C"))))
	assert.Equal(t, "GAT", DeNormalizeDNA(RevComplement(NormalizeDNA("ATC"))))
	assert.Equal(t, "ACCGGGTTTT", DeNormalizeDNA(RevComplement(NormalizeDNA("AAAACCCGGT"))))
	assert.Equal(t, "GACACAA", DeNormalizeDNA(RevComplement(NormalizeDNA("TTGTGTC"))))
}

func TestDeNormalizeDNA(t *testing.T) {
	assert.Equal(t, "A", DeNormalizeDNA(NormalizeDNA("A")))
	assert.Equal(t, "C", DeNormalizeDNA(NormalizeDNA("C")))
	assert.Equal(t, "G", DeNormalizeDNA(NormalizeDNA("G")))
	assert.Equal(t, "T", DeNormalizeDNA(NormalizeDNA("T")))
	assert.Equal(t, "ACTG", DeNormalizeDNA(NormalizeDNA("ACTG")))
}

func TestIndexToPatternStr(t *testing.T) {
	assert.Equal(t, "A", IndexToPatternStr(1, 0))
	assert.Equal(t, "AA", IndexToPatternStr(2, 0))
	assert.Equal(t, "AGT", IndexToPatternStr(3, 11))
	assert.Equal(t, "CAT", IndexToPatternStr(3, 19))
	assert.Equal(t, "CCG", IndexToPatternStr(3, 22))
	assert.Equal(t, "CCCATTC", IndexToPatternStr(7, 5437))
	assert.Equal(t, "ACCCATTC", IndexToPatternStr(8, 5437))
	assert.Equal(t, "AAAACTCGACA", IndexToPatternStr(11, 7556))
}

func TestPatternToIndex(t *testing.T) {
	assert.Equal(t, 0, PatternToIndexStr("A"))
	assert.Equal(t, 1, PatternToIndexStr("C"))
	assert.Equal(t, 2, PatternToIndexStr("G"))
	assert.Equal(t, 3, PatternToIndexStr("T"))
	assert.Equal(t, 0, PatternToIndexStr("AA"))
	assert.Equal(t, 0, PatternToIndexStr("AAAAA"))
	assert.Equal(t, 1, PatternToIndexStr("AC"))
	assert.Equal(t, 11, PatternToIndexStr("AGT"))
	assert.Equal(t, 22, PatternToIndexStr("CCG"))
	assert.Equal(t, 912, PatternToIndexStr("ATGCAA"))
	assert.Equal(t, 772508769, PatternToIndexStr("GTGAAGTGATACGAC"))
}
