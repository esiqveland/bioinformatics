package main

import (
	"strconv"
	"strings"
)

type Index struct {
	k     int
	freqs []int
}

func NewIndexStr(dna string, k int) *Index {
	normDna := NormalizeDNA(dna)
	return NewIndex(normDna, k)
}

func NewIndex(normDNA []byte, k int) *Index {

	idx := Index{
		k:     k,
		freqs: computeFrequencies(k, normDNA, createKmerArray(k)),
	}
	return &idx
}

func (idx *Index) Frequencies() []int {
	return idx.freqs
}

func (idx *Index) Results(times int) map[string]FreqWordResult {
	data := map[string]FreqWordResult{}
	for i := 0; i < len(idx.freqs); i++ {
		count := idx.freqs[i]
		if i >= times {
			dnaStr := IndexToPatternStr(idx.k, i)
			data[dnaStr] = FreqWordResult{
				Pattern: dnaStr,
				Count:   count,
			}
		}
	}

	return data
}

func computeFrequencies(k int, normDNA []byte, kmerArray []int) []int {
	dnaLen := len(normDNA)

	for i := 0; i <= dnaLen-k; i++ {
		pattern := normDNA[i : i+k]
		kmerIdx := PatternToIndex(pattern)
		kmerArray[kmerIdx] = kmerArray[kmerIdx] + 1
	}
	return kmerArray
}

func createKmerArray(k int) []int {
	num := PowInt(4, k)
	kmers := make([]int, num, num)

	return kmers
}

func FasterFrequentWordsStr(dna string, k int) map[string]FreqWordResult {
	return FasterFrequentWords(NormalizeDNA(dna), k)
}

func FasterFrequentWords(normDNA []byte, k int) map[string]FreqWordResult {
	frequentPatterns := map[string]FreqWordResult{}

	idx := NewIndex(normDNA, k)
	freqs := idx.Frequencies()

	maxCount := Maximum(freqs)
	for i := range freqs {
		if freqs[i] == maxCount {
			pattern := IndexToPatternStr(k, i)
			frequentPatterns[pattern] = FreqWordResult{
				Count:   freqs[i],
				Pattern: pattern,
			}
		}
	}

	return frequentPatterns
}

/*
   ClumpFinding(Genome, k, t, L)
       FrequentPatterns ← an empty set
       for i ← 0 to 4k − 1
           Clump(i) ← 0
       for i ← 0 to |Genome| − L
           Text ← the string of length L starting at position i in Genome
           FrequencyArray ← ComputingFrequencies(Text, k)
           for index ← 0 to 4k − 1
               if FrequencyArray(index) ≥ t
                   Clump(index) ← 1
       for i ← 0 to 4k − 1
           if Clump(i) = 1
               Pattern ← NumberToPattern(i, k)
               add Pattern to the set FrequentPatterns
       return FrequentPatterns
*/
func MovingWindowFrequentWordsFaster(dna string, kMer, windowLength, times int) map[string]FreqWordResult {
	normDNA := NormalizeDNA(dna)
	dnaLen := len(normDNA)

	wanted := map[string]FreqWordResult{}

	totalKmer := PowInt(4, kMer)
	allCounts := make([]int, totalKmer, totalKmer)

	idx := NewIndex(normDNA[0:windowLength], kMer)

	freqs := idx.Frequencies()

	for j := range freqs {
		if freqs[j] >= times {
			allCounts[j] = freqs[j]
		}
	}

	for i := 1; i <= dnaLen-windowLength-1; i++ {
		// remove the window we are moving out of
		firstPat := normDNA[i-1 : i-1+kMer]
		firstIdx := PatternToIndex(firstPat)
		freqs[firstIdx] = freqs[firstIdx] - 1

		// add the window we are moving into
		lastPat := normDNA[i+windowLength-kMer : i+windowLength]
		lastIdx := PatternToIndex(lastPat)

		// update frequency for the idx we just moved into
		freqs[lastIdx] = freqs[lastIdx] + 1

		if freqs[lastIdx] >= times {
			if freqs[lastIdx] > allCounts[lastIdx] {
				allCounts[lastIdx] = freqs[lastIdx]
			}
		}
	}

	for i := range allCounts {
		if allCounts[i] >= 1 {
			pattern := IndexToPatternStr(kMer, i)
			wanted[pattern] = FreqWordResult{
				Pattern: pattern,
				Count:   allCounts[i],
			}
		}
	}

	return wanted
}

var patToIndex = []byte{
	'A': 0,
	'C': 1,
	'G': 2,
	'T': 3,
}

// NormalizesDNA takes a dna string and returns a byte-slice, with ACTG normalized to byte values 0-3
func NormalizeDNA(dna string) []byte {
	upper := []byte(strings.ToUpper(dna))
	dnaLen := len(upper)
	buf := make([]byte, dnaLen, dnaLen)
	for i := 0; i < dnaLen; i++ {
		buf[i] = patToIndex[upper[i]]
	}
	return buf
}

func NormalizeListDNA(dna []string) sequences {
	norm := make(sequences, len(dna), len(dna))

	for i := 0; i < len(dna); i++ {
		norm[i] = NormalizeDNA(dna[i])
	}

	return norm
}

// DeNormalizeDNA takes a byteslice of normalized DNA and returns it as upper case ACTG string
func DeNormalizeDNA(dnaData []byte) string {
	dnaLen := len(dnaData)
	buf := make([]byte, dnaLen, dnaLen)
	for i := 0; i < dnaLen; i++ {
		buf[i] = indexToLetter[dnaData[i]]
	}
	return string(buf)
}

func PatternToIndexStr(pattern string) int {
	return PatternToIndex(NormalizeDNA(pattern))
}

func PatternToIndex(pattern []byte) int {
	patLen := len(pattern)

	// TODO: maybe panic here?
	if patLen == 0 {
		return 0
	} else {
		pow := patLen - 1
		sum := 0
		for i := 0; i < patLen; i++ {
			sum += int(pattern[i]) * int(Pow4(pow-i))
		}

		return sum
	}
}

var indexToLetter = []byte{
	0:   'A', // need this to fill in for the 0-padding
	1:   'C',
	2:   'G',
	3:   'T',
	'0': 'A',
	'1': 'C',
	'2': 'G',
	'3': 'T',
}

func IndexToPatternStr(k, index int) string {
	res := IndexToPattern(k, index)
	resLen := len(res)
	buf := make([]byte, resLen, resLen)
	for i := 0; i < resLen; i++ {
		buf[i] = indexToLetter[res[i]]
	}
	return string(buf)
}

func IndexToPattern(k, index int) sequence {
	// convert to base4
	data := []byte(strconv.FormatInt(int64(index), 4))

	if len(data) == k {
		return data
	}

	// add padding for K-mer
	buf := make([]byte, k, k)
	for i := k - 1; i >= 0; i-- {
		copy(buf[k-len(data):], data)
	}
	return buf
}

var complements = map[byte]byte{
	'A': 'T',
	'T': 'A',
	'G': 'C',
	'C': 'G',
}

func RevComplementStr(dna string) string {
	size := len(dna)
	buf := make([]byte, size)

	for i := 0; i < size; i++ {
		complement := complements[dna[size-i-1]]
		buf[i] = complement
	}
	return string(buf)
}

var revComplements = []byte{
	// A -> T
	0: 3,
	// T -> A
	3: 0,
	// G -> C
	2: 1,
	// C -> G
	1: 2,
}

func RevComplement(dna []byte) []byte {
	size := len(dna)
	buf := make([]byte, size)

	for i := 0; i < size; i++ {
		buf[i] = revComplements[dna[size-i-1]]
	}
	return buf
}
