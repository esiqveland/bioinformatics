package main

// Integer power: compute a**b using binary powering algorithm
// See Donald Knuth, The Art of Computer Programming, Volume 2, Section 4.6.3
func PowInt(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// Returns the nth power of 4.
func Pow4(n int) uint {
	return uint(1) << (2 * uint(n))
}

// Modular integer power: compute a**b mod m using binary powering algorithm
func PowMod(a, b, m int) int {
	a = a % m
	p := 1 % m
	for b > 0 {
		if b&1 != 0 {
			p = (p * a) % m
		}
		b >>= 1
		a = (a * a) % m
	}
	return p
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Maximum(numbers []int) int {
	max := 0
	for idx := range numbers {
		if numbers[idx] > max {
			max = numbers[idx]
		}
	}
	return max
}

func Minimum(numbers []int) (positions []int, val int) {
	min := numbers[0]

	pos := make([]int, 0, 5)
	for idx := range numbers {
		if numbers[idx] < min {
			min = numbers[idx]
		}
	}
	for idx := range numbers {
		if numbers[idx] == min {
			pos = append(pos, idx)
		}
	}
	return pos, min
}
