package main

import (
	"fmt"
	"log"
)

func HammingDistanceStr(seq1, seq2 string) int {
	if len(seq1) != len(seq2) {
		panic(fmt.Sprintf("seq1 is len %v seq2 has len", len(seq1), len(seq2)))
	}
	sum := 0
	for i := range seq1 {
		if seq1[i] != seq2[i] {
			sum = sum + 1
		}
	}

	return sum
}

func HammingDistance(seq1, seq2 sequence) int {
	if len(seq1) != len(seq2) {
		panic(fmt.Sprintf("seq1 is len %v seq2 has len", len(seq1), len(seq2)))
	}
	sum := 0
	for i := range seq1 {
		if seq1[i] != seq2[i] {
			sum = sum + 1
		}
	}

	return sum
}

func NeighborsStr(pattern string, distance int) []string {
	result := make([]string, 0, 5)
	for _, seq := range Neighbors(NormalizeDNA(pattern), distance) {
		result = append(result, DeNormalizeDNA(seq))
	}
	return result
}

type sequence []byte

type sequences []sequence

func (s sequences) Denormalize() []string {
	result := make([]string, 0, len(s))
	for i, _ := range s {
		result = append(result, DeNormalizeDNA(s[i]))
	}
	return result
}

/*
 *    IterativeNeighbors(Pattern, d)
        Neighborhood ← set consisting of single string Pattern
        for j = 1 to d
            for each string Pattern’ in Neighborhood
                add ImmediateNeighbors(Pattern') to Neighborhood
                remove duplicates from Neighborhood
        return Neighborhood
*/

func IterativeNeighbors(pattern sequence, d int) []sequence {
	neighborhood := map[string]sequence{
		string(pattern): pattern,
	}
	alreadyDone := map[string]bool{}

	for j := 1; j <= d; j++ {
		currentHood := map[string]sequence{}
		for pat := range neighborhood {
			_, ok := alreadyDone[pat]
			if ok {
				continue
			}
			alreadyDone[pat] = true

			neighbors := ImmediateNeighbors(sequence(pat))
			for _, neighbor := range neighbors {
				currentHood[string(neighbor)] = neighbor
			}
		}

	}

	all := make([]sequence, len(neighborhood))
	i := 0
	for key := range neighborhood {
		all[i] = sequence(key)
		i++
	}

	return all
}

func ImmediateNeighbors(pattern sequence) []sequence {
	neighborhood := []sequence{pattern}

	patLen := len(pattern)
	for i := 0; i < patLen; i++ {
		symbol := pattern[i]

		for n := byte(1); n < 4; n++ {
			neighbor := make([]byte, patLen)
			copy(neighbor, pattern)
			neighbor[i] = (symbol + n) % 4
			neighborhood = append(neighborhood, neighbor)
		}
	}
	return neighborhood
}

func NeighborsSimple(pattern sequence, distance int) []sequence {
	dnaLen := len(pattern)
	if distance == 0 {
		return []sequence{pattern}
	}
	if dnaLen == 1 {
		return []sequence{{0}, {1}, {2}, {3}}
	}
	neighborHood := make([]sequence, 0, 5)
	suffix := pattern[1:]
	suffixNeighbors := NeighborsSimple(suffix, distance)

	for _, text := range suffixNeighbors {
		if HammingDistance(suffix, text) < distance {
			for n := byte(0); n < 4; n++ {
				pat := append([]byte{n}, text...)
				neighborHood = append(neighborHood, pat)
			}
		} else {
			pat := append([]byte{pattern[0]}, text...)
			neighborHood = append(neighborHood, pat)
		}
	}
	return neighborHood
}

func Neighbors(dna sequence, distance int) []sequence {
	dist := uint(distance)
	dnaLen := uint(len(dna))
	numResult := dnaLen * uint(PowInt(3, distance))
	// add one for the input sequence
	alternatives := make([]sequence, numResult+2, numResult+2)
	// add the sequence itself
	alternatives[numResult] = dna

	resultSize := dnaLen * (numResult + 2)

	results := make([]byte, 2*resultSize, 2*resultSize)
	log.Printf("numResult=%v resultSize=%v\n", numResult, resultSize)

	for i := uint(0); i < resultSize; i = i + (3 * dist * dnaLen) {
		for d := uint(0); d <= dist; d++ {

			for bp := uint(0); bp < 4; bp++ {
				start := d*i + (dnaLen * bp)
				stop := start + dnaLen
				res := start / dnaLen

				log.Printf("start=%v stop=%v res=%v\n", start, stop, res)
				current := results[start:stop]
				for j := uint(0); j < d; j++ {
					current[j] = current[j] + byte(bp) + 1
					current[j] = current[j] % 4

				}

				alternatives[res] = current
			}
		}
	}

	return alternatives
}

func ApproximateSubStringCount(text, pattern string, distance int) int {
	return len(NaiveApproximateSubString(text, pattern, distance))
}

func NaiveApproximateSubString(text, pattern string, distance int) []int {
	positions := make([]int, 0, 5)
	patLen := len(pattern)
	textLen := len(text)

	for i := 0; i < textLen-patLen+1; i++ {
		window := text[i : i+patLen]
		if HammingDistanceStr(window, pattern) <= distance {
			positions = append(positions, i)
		}
	}

	return positions
}
