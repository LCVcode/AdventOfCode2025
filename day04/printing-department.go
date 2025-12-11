package day04

import (
	"aoc2025/utils"
	"iter"
	"strings"
	"fmt"
)


type Coord struct {COL, ROW int}

func (c Coord) Legal(grid [][]string) bool { return c.COL >= 0 && c.ROW >= 0 && c.COL < len(grid) && c.ROW < len(grid[0]) }


func GridCoords(width, height int) iter.Seq[Coord] {
	return func(yield func(Coord) bool) {
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if !yield(Coord{y, x}) { return }
			}
		}
	}
}


func Neighbors(c Coord) iter.Seq[Coord] {
	return func(yield func(Coord) bool) {
		for v := -1; v < 2; v++ {
			for u := -1; u < 2; u++ {
				if v == 0 && u == 0 { continue }
				if !yield(Coord{c.COL+u, c.ROW+v}) { return }
			}
		}
	}
}


func LegalNeighbors(c Coord, maxW, maxH int) iter.Seq[Coord] {
	return func(yield func(Coord) bool) {
		for n := range Neighbors(c) {
			if n.COL < 0 || n.ROW < 0 || n.COL >= maxH || n.ROW >= maxW { continue }
			if !yield(n) { return }
		}
	}
}


func Part1() int {
	inputs, err := utils.ReadFileLines("day04/inputs.txt")
	if err != nil {
		panic(err)
	}

	height := len(inputs)
	width := len(inputs[0])
	total := 0

	for c := range GridCoords(width, height) {
		if inputs[c.COL][c.ROW:c.ROW+1]=="." { continue }
		local := 0
		for n := range LegalNeighbors(c, width, height) {
			// if inputs[n.COL][n.ROW:n.ROW+1]=="@" { local++ }
			if inputs[n.COL][n.ROW:n.ROW+1]=="@" { local++ }
		}
		if local < 4 { 
			total++ 
		}
	}

	return total
}


func gridGet(grid [][]string, c Coord) (string, error) {
	if !c.Legal(grid) {
		return "", fmt.Errorf("invalid coordinate: %v", c)
	}
	result := grid[c.COL][c.ROW]
	return result, nil
}


func Part2() int {
	inputs, err := utils.ReadFileLines("day04/inputs.txt")
	if err != nil {
		panic(err)
	}

	height := len(inputs)
	width := len(inputs[0])
	grid := make([][]string, len(inputs))
	for i, line := range inputs {
		grid[i] = strings.Split(line, "")
	}
	total := 0
	changed := []Coord{Coord{0, 0}}

	for len(changed) > 0 {
		changed = changed[:0]
		for c := range GridCoords(width, height) {
			cell, _ := gridGet(grid, c)
			if cell == "." { continue }
			local := 0
			for n := range LegalNeighbors(c, width, height) {
				cell, err := gridGet(grid, n) 
				if err != nil { continue }
				if cell == "." { continue }
				local++
			}
			if local < 4 {
				changed = append(changed, c)
				continue
			}
		}
		for _, c := range changed {
			grid[c.COL][c.ROW] = "."
		}
		total += len(changed)

	}

	return total
}
