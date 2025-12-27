package main


import (
	"fmt"
	"time"
	// "aoc2025/day01"
	// "aoc2025/day02"
	// "aoc2025/day03"
	// "aoc2025/day04"
	"aoc2025/day05"
	"aoc2025/day06"
)


func main() {
	start := time.Now()
	/*
	d1p1Result := day01.Part1()
	fmt.Printf("Day 1 Part 1 | %d | %v\n", d1p1Result, time.Since(start))
	start = time.Now()
	d1p2Result := day01.Part2()
	fmt.Printf("Day 1 Part 2 | %d | %v\n", d1p2Result, time.Since(start))
	start = time.Now()
	d2p1Result := day02.Part1()
	fmt.Printf("Day 2 Part 1 | %d | %v\n", d2p1Result, time.Since(start))
	start = time.Now()
	d2p2Result := day02.Part2()
	fmt.Printf("Day 2 Part 2 | %d | %v\n", d2p2Result, time.Since(start))
	start = time.Now()
	d3p1Result := day03.Part1()
	fmt.Printf("Day 3 Part 1 | %d | %v\n", d3p1Result, time.Since(start))
	d3p2Result := day03.Part2()
	start = time.Now()
	fmt.Printf("Day 3 Part 2 | %d | %v\n", d3p2Result, time.Since(start))
	start = time.Now()
	d4p1Result := day04.Part1()
	fmt.Printf("Day 4 Part 1 | %d | %v\n", d4p1Result, time.Since(start))
	d4p2Result := day04.Part2()
	start = time.Now()
	fmt.Printf("Day 4 Part 2 | %d | %v\n", d4p2Result, time.Since(start))
	*/
	start = time.Now()
	d5p1Result := day05.Part1()
	fmt.Printf("Day 5 Part 1 | %d | %v\n", d5p1Result, time.Since(start))
	start = time.Now()
	d5p2Result := day05.Part2()
	fmt.Printf("Day 5 Part 2 | %d | %v\n", d5p2Result, time.Since(start))
	start = time.Now()
	d6p1Result := day06.Part1()
	fmt.Printf("Day 6 Part 1 | %d | %v\n", d6p1Result, time.Since(start))
	start = time.Now()
	d6p2Result := day06.Part2()
	fmt.Printf("Day 6 Part 2 | %d | %v\n", d6p2Result, time.Since(start))
}
