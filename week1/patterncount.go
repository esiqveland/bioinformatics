package main

import (
	"fmt"
	"unicode/utf8"
)

func SubStringPositionsAsString(dna, pattern string) string {
	positions := SubStringPositions(dna, pattern)

	str := ""
	for _, position := range positions {
		str += fmt.Sprintf("%d ", position)
	}
	if len(positions) > 0 {
		// remove the last trailing space
		return str[:len(str)-1]
	} else {
		return str
	}
}

func SubStringPositions(dna, pattern string) []int {
	positions := []int{}
	lenText := utf8.RuneCountInString(dna)
	lenPattern := utf8.RuneCountInString(pattern)

	for i := 0; i <= lenText-lenPattern; i++ {
		section := dna[i : i+lenPattern]
		if section == pattern {
			positions = append(positions, i)
		}
	}

	return positions
}

func PatternCount(text, pattern string) int {
	count := 0
	lenText := utf8.RuneCountInString(text)
	lenPattern := utf8.RuneCountInString(pattern)

	for i := 0; i <= lenText-lenPattern; i++ {
		section := text[i : i+lenPattern]
		//fmt.Println("i=%v lenPattern=%v section=%v", i, i+lenPattern, section)

		if section == pattern {
			count = count + 1
		}
	}
	return count
}

type FreqWordResult struct {
	RevPattern string
	Pattern    string
	RevCount   int
	Count      int
}

func (f *FreqWordResult) String() string {
	return fmt.Sprintf("Pattern: %v Count: %v RevCount: %v Total: %v", f.Pattern, f.Count, f.RevCount, f.Count+f.RevCount)
}

/*
   FrequentWords(Text, k)
       FrequentPatterns ← an empty set
       for i ← 0 to |Text| − k
           Pattern ← the k-mer Text(i, k)
           Count(i) ← PatternCount(Text, Pattern)
       maxCount ← maximum value in array Count
       for i ← 0 to |Text| − k
           if Count(i) = maxCount
               add Text(i, k) to FrequentPatterns
       remove duplicates from FrequentPatterns
       return FrequentPatterns

*/
func FrequentWords(text string, k int) map[string]FreqWordResult {
	freqPatterns := map[string]int{}
	lenText := utf8.RuneCountInString(text)
	for i := 0; i <= lenText-k; i++ {
		pattern := text[i : i+k]
		if _, ok := freqPatterns[pattern]; !ok {
			count := PatternCount(text, pattern)
			freqPatterns[pattern] = count
		}
	}

	max := -1
	for _, count := range freqPatterns {
		if count > max {
			max = count
		}
	}

	results := map[string]FreqWordResult{}
	for pattern, count := range freqPatterns {
		if count == max {
			revPattern := RevComplementStr(pattern)
			data := FreqWordResult{
				Pattern:    pattern,
				RevPattern: revPattern,
				Count:      count,
				RevCount:   0,
			}

			if revCount, ok := freqPatterns[revPattern]; ok {
				//log.Printf("Found complement: %v", revPattern)
				data.RevCount = revCount
			}

			results[pattern] = data
		}
	}
	return results
}

func FrequentKWords(text string, kMerFrom, kMerTo int) map[int]map[string]FreqWordResult {
	results := map[int]map[string]FreqWordResult{}

	for k := kMerFrom; k <= kMerTo; k++ {
		data := FrequentWords(text, k)

		resultList := map[string]FreqWordResult{}
		for _, val := range data {
			resultList[val.Pattern] = val
		}
		results[k] = resultList
	}

	return results
}

// Clump Finding Problem: Find patterns forming clumps in a string.
func MovingWindowFrequentWords(dna string, kMer, windowLength, times int) map[string]FreqWordResult {
	dnaLen := len(dna)

	wanted := map[string]FreqWordResult{}

	for i := 0; i <= dnaLen-windowLength; i = i + 1 {
		stop := Min(i+windowLength, dnaLen-1)
		start := Min(i, stop-windowLength)

		dnaWindow := dna[start:stop]

		results := FrequentWords(dnaWindow, kMer)

		for pattern, result := range results {
			if result.Count >= times {
				stored, ok := wanted[pattern]
				if ok && stored.Count > result.Count {
					//wanted[pattern] = result
					// we already have a better result, skip this
					continue
				} else {
					wanted[pattern] = result
				}
			}
		}
	}

	return wanted
}

func main() {
	//pattern := "TGTGGAATG"
	//count := PatternCount(text, pattern)
	//fmt.Printf("pattern=%v count=%v\n", pattern, count)
	results := FrequentWords(text2, 11)
	fmt.Printf("%+v", results)
}

const text = "GTTGTGGAATCTGTGGAAATGTGGAATTGTGGAATCCTGTGGAAGTGTGGAAGGGTGTGGAATAGTATTCTTGGGAGTCATGTGGAAGATGTGGAACTGTGGAATTGTGGAAGAGTGTGGAACTGCTGTGGAATTTAGTTGTGGAACTGTGGAATATGTGGAAAGCGTGTGGAATGTGGAAGTAGCAATCTGTGGAATCTTATGTGGAACTCTATGTGGAATGTGGAATTGTGGAATGTGGAAATGTATGTGGAATGTGGAACTGTGGAATGTGGAAAATCTGTTGTGGAAACACTGTGGAAATTGGTTGTGGAACATCATGTGGAATGTGGAAATTGTGGAACGGACTTGTGGAAAATGTGGAATGTGGAATGTGGAACCGTGTGGAATAGTTTGTGGAACTGTGGAATGTGGAATGTGGAATGTGGAACACACTGCTGTGGAACTGTGGAAATGTGGAAACGAGGTGTGGAATCTATGTGGAAATGTGGAATAGATGACTGTGGAAGCGTGTGGAATTTTTGTGGAATGTGGAACTGTGGAATCGCAATGCTGTGGAATGTGGAATGTGGAAGTGTGGAACATGTGGAACATGTGGAATGTGGAACTGTGGAACCCTTGTGGAATAAACATGTGGAAGGTATTGTGGAACTGTGTGGAACATGTGGAATGTGGAATTGTGTGGAATCTTGTGGAATGTGGAATGTGGAAATAATACTGTGGAAATGTGGAATTTCACATGTGGAATGTGGAATCGTGTGGAAGCTTCGAAGTGTGACTGTGGAAGTGTGGAACCGGGCGTGTGGAAGGGGTGTGGAAGTGTGGAAGTTAATTGTGGAAGTGTGGAAAGCAACACGGTGTGGAATGTGGAATGTGGAACTGTGGAATGAACTGTGGAATCTGTATCTCCTTGTGGAAGTGTGGAATGTGGAAATGTGGAATGTGGAACCGTAGGTGTGGAATGTGGAACTGTGGAA"
const text2 = "GTTTGGCGGGCCTCGGGCCTCGGCAGGCGAGGCAGGCGACGGTTCAGGGCCTCCGCTTGAGACGGTTCAGGGCCTCCGCTTGAGAGTTTGGCGGGCCTCCGGTTCAGGCAGGCGACGCTTGAGACGCTTGAGACGGTTCACGCTTGAGACGGTTCACGCTTGAGACGCTTGAGACGCTTGAGAGGGCCTCGGGCCTCGTTTGGCCGGTTCAGGGCCTCCGCTTGAGAGTTTGGCGGGCCTCCGCTTGAGAGGGCCTCGTTTGGCCGCTTGAGACGCTTGAGAGTTTGGCGGCAGGCGAGGCAGGCGACGGTTCACGGTTCAGTTTGGCCGCTTGAGACGCTTGAGACGGTTCAGTTTGGCGTTTGGCCGCTTGAGACGCTTGAGAGGGCCTCGGGCCTCGGGCCTCGTTTGGCCGGTTCAGGGCCTCGGGCCTCGGCAGGCGACGGTTCACGCTTGAGAGGGCCTCGGCAGGCGACGCTTGAGAGTTTGGCCGGTTCACGCTTGAGAGGCAGGCGACGCTTGAGAGGCAGGCGACGCTTGAGACGGTTCAGTTTGGCCGCTTGAGAGGGCCTCCGCTTGAGAGTTTGGCGTTTGGCGGCAGGCGAGGGCCTCCGCTTGAGAGTTTGGCCGCTTGAGAGTTTGGCGGCAGGCGAGGCAGGCGAGGCAGGCGACGGTTCACGGTTCAGGCAGGCGAGTTTGGCGGCAGGCGACGGTTCACGGTTCAGGCAGGCGACGCTTGAGAGTTTGGCGGCAGGCGACGGTTCAGTTTGGCGGGCCTCCGGTTCACGCTTGAGACGGTTCAGTTTGGCGTTTGGC"
