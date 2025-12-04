package day03

import (
	"aoc2025/utils"
	"strconv"
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

func Part1() int {
	raw, err := utils.ReadFileLines("day03/inputs.txt")
	if err != nil {
		panic(err)
	}

	var total int
	for _, bank := range raw {
		total += maxJoltage(bank)
	}

	/*
	values := make([]int, len(raw))
	for i, s := range raw {
		values[i], _ = strconv.Atoi(s)
	}
	fmt.Println(values)
	*/
	
	return total
}


func Part2() int {
	return 0
}
