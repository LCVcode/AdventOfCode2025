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

func nDupeInt(sub, n int) int {
	str := strconv.Itoa(sub)
	result, _ := strconv.Atoi(strings.Repeat(str, n))
	return result
}

func dupeInt(sub int) int {
	// Duplicates an integer (9 -> 99 or 101 -> 101101)
	return nDupeInt(sub, 2)
}

func Factorize(n int) []int {
	var facts []int
	for i := 1; i <= n/2; i++ {
		if n%i != 0 { continue }
		facts = append(facts, i)
	}
	return facts
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
		// Shortcut out if range is now undefined
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


// chunk splits s into substrings of length n.
// The last chunk may be shorter.
func chunk(s string, n int) []string {
    if n <= 0 || s == "" {
        return nil
    }
    out := make([]string, 0, (len(s)+n-1)/n)
    for len(s) > 0 {
        if len(s) < n {
            out = append(out, s)
            break
        }
        out = append(out, s[:n])
        s = s[n:]
    }
    return out
}


func doesNCycle(str string) bool {
	facts := Factorize(len(str))
	for _, fact := range facts {
		chunks := chunk(str, fact)
		allSame := true
		for _, c := range chunks[1:] {
			if c != chunks[0] {
				allSame = false
				break
			}
		}
		if allSame { return true }
	}
	return false
}


func BruteForceNCycles(lo, hi string) int {
	var total int
	ilo, _ := strconv.Atoi(lo)
	ihi, _ := strconv.Atoi(hi)
	for i := ilo; i <= ihi; i++ {
		str := strconv.Itoa(i)
		if doesNCycle(str) {
			total += i
		}

	}
	return total
}


func Part2() int {
	raw, err := utils.ReadFileLines("day02/inputs.txt")
	if err != nil {
		panic(err)
	}
	var total int
	ranges := strings.Split(raw[0], ",")

	for _, rng := range ranges {
		if len(rng) == 0 { continue }
		lo, hi, err := splitRange(rng)
		if err != nil { panic(err) }
		total += BruteForceNCycles(lo, hi)
	}

	return total
}
