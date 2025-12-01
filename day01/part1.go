package day01

import (
	"aoc2025/utils"
	"fmt"
	"strconv"
)

func Part1() int {
	inputs, err := utils.ReadFileLines("day01/inputs1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(inputs))

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
	inputs, err := utils.ReadFileLines("day01/inputs1.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(inputs))

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
		// fmt.Printf("%s\t%d\t%d\n", line, pos, total)
	}

	fmt.Printf("Position: %d\n", pos)
	return total
}
