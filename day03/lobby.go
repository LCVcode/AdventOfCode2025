package day03

import (
	"aoc2025/utils"
	"strconv"
	"strings"
)

func maxJoltage(bank string) int {
	best := "-1"
	i := -1
	l := len(bank)
	for j := range bank[:l-1] {
		joltage := bank[j:j+1]
		if joltage > best {
			best = joltage
			i = j
		}
	}
	big, _ := strconv.Atoi(best)
	offset := i + 1
	best = "-1"
	for k := range bank[offset:] {
		joltage := bank[k+offset:k+offset+1]
		if joltage > best {
			best = joltage
			i = k
		}
	}
	small, _ := strconv.Atoi(best)
	return 10*big + small
}


func maxJoltage12(bank string) int {
	result := make([]string, 12)
	l := len(bank)
	var offset int
	for i := 0; i < 12; i++ {
		best := "-1"
		index := 0
		subBank := bank[i+offset:l-11+i]
		for j := 0; j < len(subBank); j++ {
			jolt := subBank[j:j+1]
			if jolt > best {
				best = jolt
				index = j
			}
		}
		offset += index
		result[i] = best
	}
	total, _ := strconv.Atoi(strings.Join(result, ""))
	return total
}

func Part1() int {
	raw, err := utils.ReadFileLines("day03/inputs.txt")
	if err != nil {
		panic(err)
	}

	var total int
	for _, bank := range raw {
		total += maxJoltage(bank)
	}
	return total
}


func Part2() int {
	raw, err := utils.ReadFileLines("day03/inputs.txt")
	if err != nil {
		panic(err)
	}

	var total int
	for _, bank := range raw {
		total += maxJoltage12(bank)
	}
	return total
}
