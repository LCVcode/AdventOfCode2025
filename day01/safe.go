package day01

import (
	"aoc2025/utils"
	"strconv"
)

func Part1() int {
	inputs, err := utils.ReadFileLines("day01/inputs.txt")
	if err != nil {
		panic(err)
	}

	var total, pos int
	pos = 50
	for _, line := range inputs {
		dir := line[:1]
		num, _ := strconv.Atoi(line[1:])
		if dir == "L" {
			pos -= num
			for pos < 0 {
				pos += 100
			}
		} else {
			pos += num
			pos %= 100
		}
		if pos == 0 {
			total++
		}
	}

	return total
}

func Part2() int {
	inputs, err := utils.ReadFileLines("day01/inputs.txt")
	if err != nil {
		panic(err) 
	}

	var total, pos int
	pos = 50
	for _, line := range inputs {
		dir := line[:1]
		num, _ := strconv.Atoi(line[1:])
		
		if dir == "L" {
			pos -= num
		} else {
			pos += num
		}
		
		for pos < 0 {
			pos += 100
			// Handles special case where a left turn started from 0
			if pos + num != 100 {
				total += 1
			}
		}
		for pos > 99 {
			pos -= 100
			total += 1
		}
		if pos == 0 && dir == "L" {
			total += 1
		}
	}
	return total
}
