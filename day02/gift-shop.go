package day02

import (
	"aoc2025/utils"
	"strings"
	"fmt"
	"strconv"
	"math"
)

func splitRange(rng string) (string, string, error) {
	i := strings.Index(rng, "-")
	if i == -1 {
		return "", "", fmt.Errorf("no '-' found in %q", rng)
	}
	return rng[:i], rng[i+1:], nil
}

func dupeInt(sub int) int {
	// Duplicates an integer (9 -> 99 or 101 -> 101101)
	str := strconv.Itoa(sub)
	result, _ := strconv.Atoi(strings.Repeat(str, 2))
	return result
}

func Part1() int {
	raw, err := utils.ReadFileLines("day02/inputs.txt")
	if err != nil {
		panic(err)
	}
	
	var total int
	ranges := strings.Split(raw[0], ",")

	for _, rng := range ranges {
		if len(rng) == 0 { continue }
		lo, hi, err := splitRange(rng)
		if err != nil {
			panic(err)
		}

		// Limit range limits to even-digited values (odds can be ignored)
		//  lo to be rounded up   to the next even-digit number
		//  hi to be rounded down to the next even-digit number
		if len(lo) % 2 == 1 {
			lo = strconv.Itoa(int(math.Pow10(len(lo))))
		}
		if len(hi) % 2 == 1 {
			hi = strings.Repeat("9", len(hi) - 1)
		}
		if len(lo) > len(hi) { continue }

		// Limit strings are each split in half
		i := len(lo) / 2
		loMajor, _ := strconv.Atoi(lo[:i])
		loMinor, _ := strconv.Atoi(lo[i:])
		i = len(hi) / 2
		hiMajor, _ := strconv.Atoi(hi[:i])
		hiMinor, _ := strconv.Atoi(hi[i:])

		// ======================
		// This is the main logic
		// ======================

		// Special case: major values match
		if loMajor == hiMajor {
			if loMinor <= loMajor && hiMinor >= hiMajor {
				total += dupeInt(loMajor)
			}
			continue
		}

		// General case: loMajor < hiMajor
		for i := loMajor + 1; i < hiMajor; i++ {
			total += dupeInt(i)
		}
		if loMinor <= loMajor {
			total += dupeInt(loMajor)
		}
		if hiMinor >= hiMajor {
			total += dupeInt(hiMajor)
		}
	}
	return total
}

func Part2() int {
	return 0
}
