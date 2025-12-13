package day05

import (
	"aoc2025/utils"
	"strconv"
	"strings"
	"sort"
)

type IDRange struct {
	Low int
	High int
}

func (rng IDRange) Contains(x int) bool { return rng.Low <= x && x <= rng.High }

func (rng IDRange) Span() int { return rng.High - rng.Low + 1}

func Max(a, b int) int { if a > b { return a }; return b }

func Part1() int {
	inputs, err := utils.ReadFileLines("day05/inputs.txt")
	if err != nil {
		panic(err)
	}
	pivot := 0
	for i, v := range inputs {
		if v == "" {
			pivot = i
			break
		}
	}
	rangeStrings := inputs[:pivot]
	idStrings := inputs[pivot+1:]

	idRanges := make([]IDRange, len(rangeStrings))
	for i, rng := range rangeStrings {
		parts := strings.Split(rng, "-")
		low, _ := strconv.Atoi(parts[0])
		high, _ := strconv.Atoi(parts[1])
		idRanges[i] = IDRange{low, high}
	}

	total := 0
	for _, id := range idStrings {
		val, _ := strconv.Atoi(id)
		detected := false
		for _, rng := range idRanges {
			if rng.Contains(val) { 
				detected = true
				break
			}
		}

		if detected {
			total++
		}
	}

	return total
}

func Part2() int {
	inputs, err := utils.ReadFileLines("day05/inputs.txt")
	if err != nil {
		panic(err)
	}
	pivot := 0
	for i, v := range inputs {
		if v == "" {
			pivot = i
			break
		}
	}
	rangeStrings := inputs[:pivot]
	idRanges := make([]IDRange, len(rangeStrings))
	for i, rng := range rangeStrings {
		parts := strings.Split(rng, "-")
		low, _ := strconv.Atoi(parts[0])
		high, _ := strconv.Atoi(parts[1])
		idRanges[i] = IDRange{low, high}
	}

	// Sort IDRanges by the low value - simplifies future logic
	sort.Slice(idRanges, func(i, j int) bool {
		return idRanges[i].Low < idRanges[j].Low
	})

	finalRanges := make([]IDRange, 0)
	for len(idRanges) > 1 {
		ref := idRanges[0]
		idRanges = idRanges[1:]
		for len(idRanges) > 0 && idRanges[0].Low <= ref.High {
			ref.High = Max(ref.High, idRanges[0].High)
			idRanges = idRanges[1:]
		}
		finalRanges = append(finalRanges, ref)
	}

	if len(idRanges) > 0 { finalRanges = append(finalRanges, idRanges[0]) }

	total := 0
	for _, r := range finalRanges {
		total += r.Span()
	}
	return total
}
