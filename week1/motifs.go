package main

import (
	"fmt"
	"os"
)

/*
   MotifEnumeration(Dna, k, d)
       Patterns ← an empty set
       for each k-mer Pattern in the first string in Dna
           for each k-mer Pattern’ differing from Pattern by at most d mismatches
               if Pattern' appears in each string from Dna with at most d mismatches
                   add Pattern' to Patterns
       remove duplicates from Patterns
       return Patterns

*/

// MotifEnumeration returns all motifs that are present in all sequences
// with at most d mismatches
func MotifEnumeration(dna sequences, k, d int) sequences {
	if len(dna) < 2 {
		panic("must give more than 1 sequence")
	}
	patterns := map[string]sequence{}

	firstPattern := dna[0]
	for i := 0; i <= len(firstPattern)-k; i++ {
		kmer := firstPattern[i : i+k]

		neighbHood := NeighborsSimple(kmer, d)

		for n := 0; n < len(neighbHood); n++ {
			missing := false
			neib := neighbHood[n]
			for j := 1; j < len(dna); j++ {
				if !kmerExistsWithMisMatches(dna[j], neib, d) {
					missing = true
					break
				}
			}

			if !missing {
				patterns[string(neib)] = neib
			}
		}
	}

	results := sequences{}
	for _, value := range patterns {
		results = append(results, value)
	}

	return results
}

func kmerExistsWithMisMatches(text sequence, kmer sequence, distance int) bool {
	textLen := len(text)
	patLen := len(kmer)

	for i := 0; i <= textLen-patLen; i++ {
		window := text[i : i+patLen]
		if HammingDistance(window, kmer) <= distance {
			return true
		}
	}

	return false
}

func MedianString(dna sequences, k int) sequences {
	distance := k * len(dna) * len(dna[0])
	median := -1

	count := int(Pow4(k))
	for i := 0; i < count; i++ {
		kmer := IndexToPattern(k, i)

		//TODO: fix this weird bug with having to do NormalizeDNA(DeNormalizeDNA(kmer))
		dist := DistanceBetweenPatternAndStrings(dna, NormalizeDNA(DeNormalizeDNA(kmer)))
		fmt.Fprintf(os.Stderr, "dist=%v distance=%v kmer=%v\n", dist, distance, DeNormalizeDNA(kmer))
		if dist < distance {
			distance = dist
			median = i
		}
	}

	return sequences{IndexToPattern(k, median)}
}

func DistanceBetweenPatternAndStrings(dna sequences, pattern sequence) int {
	k := len(pattern)

	distance := 0
	for _, text := range dna {

		hamming := HammingDistance(pattern, text[0:k])
		for i := 0; i <= len(text)-k; i++ {
			kmer := text[i : i+k]

			kmerDistance := HammingDistance(pattern, kmer)
			if kmerDistance < hamming {
				hamming = kmerDistance
			}
		}

		distance = distance + hamming
	}

	fmt.Printf("DistanceBetweenPatternAndStrings: %v result=%v\n", DeNormalizeDNA(pattern), distance)

	return distance
}

func MostProbableKmer(dna sequence, k int, matrix ProfileMatrix) sequence {

	var bestPattern sequence
	best := 0.0
	for i := 0; i <= len(dna)-k; i++ {
		kmer := dna[i : i+k]
		score := matrix.Score(kmer)
		if score > best {
			best = score
			bestPattern = kmer
		}
	}

	return bestPattern
}

func (mat *ProfileMatrix) Get(nuc int, pos int) float64 {
	return mat.data[nuc][pos]
}

func (mat *ProfileMatrix) Score(pattern sequence) float64 {
	score := 1.0
	for i := 0; i < len(pattern); i++ {
		score = score * mat.Get(int(pattern[i]), i)
	}

	return score
}

type ProfileMatrix struct {
	data [][]float64
}
