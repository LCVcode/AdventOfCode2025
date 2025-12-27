package day06

import (
	"aoc2025/utils"
	// "fmt"
	"strings"
	"strconv"
	"regexp"
)

func convertLine(line string) []int {
	re := regexp.MustCompile(` +`)
	firstChar := line[0:1]
	for firstChar == " " {
		line = line[1:]
		firstChar = line[0:1]
	}
	line = re.ReplaceAllString(line, " ")
	count := strings.Count(line, " ") + 1
	ints := make([]int, count)
	split := strings.Split(line, " ")
	for i := 0; i < count; i++ {
		ints[i], _ = strconv.Atoi(split[i])
	}
	return ints
}

func extractOps(line string) []string {
	re := regexp.MustCompile(` +`)
	firstChar := line[0:1]
	for firstChar == " " {
		line = line[1:]
		firstChar = line[0:1]
	}
	line = re.ReplaceAllString(line, " ")
	return strings.Split(line, " ")
}

func Part1() int {
	inputs, err := utils.ReadFileLines("day06/inputs.txt")
	if err != nil {
		panic(err)
	}

	l := len(inputs)
	valueMatrix := make([][]int, l - 1)
	for i, line := range inputs[:l - 1] {
		valueMatrix[i] = convertLine(line)
	}
	ops := extractOps(inputs[l-1])
	total := 0
	for i, op := range ops {
		switch op {
		case "+":
			subTotal := 0
			for j := 0; j < len(valueMatrix); j++ {
				subTotal += valueMatrix[j][i]
			}
			total += subTotal
		case "*":
			subTotal := 1
			for j := 0; j < len(valueMatrix); j++ {
				subTotal = valueMatrix[j][i] * subTotal
			}
			total += subTotal
		}
	}
	return total
}

func Part2() int {
	_, err := utils.ReadFileLines("day06/inputs.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(inputs)
	return 0
}
