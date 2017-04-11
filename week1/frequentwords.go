package main

import (
	"bytes"
	"fmt"
	"log"
	"sort"
)

/*
  FrequentWordsWithMismatches(Text, k, d)
        FrequentPatterns ← an empty set
        Neighborhoods ← an empty list
        for i ← 0 to |Text| − k
            add Neighbors(Text(i, k), d) to Neighborhoods
        form an array NeighborhoodArray holding all strings in Neighborhoods
        for i ← 0 to |NeighborhoodArray| − 1
            Pattern ← NeighborhoodArray(i)
            Index(i) ← PatternToNumber(Pattern)
            Count(i) ← 1
        SortedIndex ← Sort(Index)
        for i ← 0 to |NeighborhoodArray| − 1
            if SortedIndex(i) = SortedIndex(i + 1)
                Count(i + 1) ← Count(i) + 1
       maxCount ← maximum value in array Count
       for i ← 0 to |NeighborhoodArray| − 1
           if Count(i) = maxCount
               Pattern ← NumberToPattern(SortedIndex(i), k)
               add Pattern to FrequentPatterns
       return FrequentPatterns
*/

func FrequentWordsWithMismatches(text sequence, k int, d int) []sequence {
	freqPatterns := make([]sequence, 0, 5)
	neighborhoods := []sequence{}

	for i := 0; i < len(text)-k; i++ {
		subStr := text[i : i+k]
		//neighbors := IterativeNeighbors(subStr, d)
		neighbors := NeighborsSimple(subStr, d)
		neighborhoods = append(neighborhoods, neighbors...)
	}

	counts := make([]int, len(neighborhoods))
	index := make([]int, len(neighborhoods))
	for i := range neighborhoods {
		idx := PatternToIndex(neighborhoods[i])
		index[i] = idx
		counts[i] = 1
	}

	sort.Ints(index)
	for i := 0; i < len(neighborhoods)-1; i++ {
		if index[i] == index[i+1] {
			counts[i+1] = counts[i] + 1
		}
	}

	maxCount := Maximum(counts)
	for i := 0; i < len(neighborhoods)-1; i++ {
		if counts[i] == maxCount {
			pattern := IndexToPattern(k, index[i])
			fmt.Printf("Count: %v Pattern: %v", maxCount, pattern)
			log.Printf("Count: %v Pattern: %v", maxCount, DeNormalizeDNA(pattern))

			freqPatterns = append(freqPatterns, pattern)
		}
	}

	return freqPatterns
}

func FrequentWordsWithMismatchesRevComplement(text sequence, k int, d int) []sequence {
	freqPatterns := make([]sequence, 0, 5)
	neighborhoods := []sequence{}

	for i := 0; i < len(text)-k; i++ {
		subStr := text[i : i+k]
		//neighbors := IterativeNeighbors(subStr, d)
		neighbors := NeighborsSimple(subStr, d)
		neighborhoods = append(neighborhoods, neighbors...)
	}

	counts := make([]int, len(neighborhoods))
	index := make([]int, len(neighborhoods))
	patternsSeen := map[int]int{}
	for i := range neighborhoods {
		idx := PatternToIndex(neighborhoods[i])
		index[i] = idx
		counts[i] = 1
		patternsSeen[idx] = i
	}

	indexSorted := make([]int, len(index))
	copy(indexSorted, index)
	sort.Ints(indexSorted)
	for i := 0; i < len(neighborhoods)-1; i++ {
		if indexSorted[i] == indexSorted[i+1] {
			counts[i+1] = counts[i] + 1
		}
	}

	done := map[string]bool{}
	for i := 0; i < len(neighborhoods); i++ {
		dna := string(neighborhoods[i])
		_, ok := done[dna]
		if ok {
			continue
		} else {
			done[dna] = true
		}

		// find the reverse complement
		// if we have seen the rev complement, add it's count to our current pattern
		rev := RevComplement(neighborhoods[i])
		if bytes.Equal(neighborhoods[i], rev) {
			fmt.Println("Skipping adding palindrome to itself")
			continue
		}
		revIdx := PatternToIndex(rev)
		if patternsSeen[revIdx] > 0 {
			counts[i] = counts[i] + counts[patternsSeen[revIdx]]
		}
	}

	maxCount := Maximum(counts)
	fmt.Printf("maxCount=%v \n", maxCount)
	for i := 0; i < len(neighborhoods)-1; i++ {
		if counts[i] == maxCount {
			pattern := IndexToPattern(k, indexSorted[i])
			freqPatterns = append(freqPatterns, pattern)
		}
	}

	return freqPatterns
}
