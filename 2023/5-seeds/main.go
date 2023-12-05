package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"math"
	"slices"
	"strings"
)

type mapping struct {
	SourceStart      int
	DestinationStart int
	Length           int
}

func main() {
	maps := [...][]mapping{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}

	file := util.FileFromArgs()
	b := make([]byte, 1)
	line := ""
	seeds := []int{}
	mapPointer := -1

	// Parse mappings
	for {
		_, err := file.Read(b)
		if err == io.EOF || b[0] == '\n' {
			if strings.HasPrefix(line, "seeds: ") && len(seeds) == 0 {
				seeds = util.StringToInts(line[7:], " ")

				line = ""
				continue
			}

			if strings.TrimSpace(line) != "" {
				if strings.HasSuffix(line, "map:") {
					if mapPointer > -1 {
						maps[mapPointer] = fillEmptySpace(maps[mapPointer])
					}
					mapPointer++
					line = ""
					continue
				}

				parts := util.StringToInts(line, " ")
				maps[mapPointer] = append(maps[mapPointer], mapping{
					SourceStart:      parts[1],
					DestinationStart: parts[0],
					Length:           parts[2],
				})
			}

			if err == io.EOF {
				break
			}

			line = ""
			continue
		}

		line += string(b[0])
	}

	// Part 1 - lowest seed number
	lowestSeedNumber := -1
	for _, seed := range seeds {
		x := seed
		for i := 0; i < len(maps); i++ {
			for _, mapping := range maps[i] {
				if x >= mapping.SourceStart && x < mapping.SourceStart+mapping.Length {
					x = mapping.DestinationStart + (x - mapping.SourceStart)
					break
				}
			}

		}

		if lowestSeedNumber == -1 || x < lowestSeedNumber {
			lowestSeedNumber = x
		}
	}

	// Part 2 - lowest seed number from range
	lowestSeedNumberFromRange := -1
	seedPair := chunkSlice(seeds)

	jobs := make(chan []int, len(seedPair))
	results := make(chan int, len(seedPair))

	numberOfWorkers := math.Min(float64(len(seedPair)), 12)
	for i := 0; i < int(numberOfWorkers); i++ {
		go findSeedLocationWorker(maps, jobs, results)
	}

	for _, seed := range seedPair {
		jobs <- seed
	}
	close(jobs)

	var x int
	for a := 0; a < len(seedPair); a++ {
		x = <-results
		if lowestSeedNumberFromRange == -1 || x < lowestSeedNumberFromRange {
			lowestSeedNumberFromRange = x
		}
	}

	fmt.Println(lowestSeedNumber)
	fmt.Println(lowestSeedNumberFromRange)
}

func chunkSlice(slice []int) [][]int {
	chunks := [][]int{}
	for i := 0; i < len(slice); i += 2 {
		chunks = append(chunks, slice[i:i+2])
	}

	return chunks
}

func fillEmptySpace(mapp []mapping) []mapping {
	slices.SortFunc(mapp, func(a, b mapping) int {
		return a.SourceStart - b.SourceStart
	})

	newMap := []mapping{}
	start := 0
	for _, rng := range mapp {
		if rng.SourceStart > start {
			newMap = append(newMap, mapping{
				SourceStart:      start,
				DestinationStart: start,
				Length:           rng.SourceStart - start,
			})
		}

		newMap = append(newMap, rng)

		start = rng.SourceStart + rng.Length
	}

	return newMap
}

func findSeedLocationWorker(maps [7][]mapping, seeds <-chan []int, location chan<- int) {
	for seed := range seeds {
		lowest := -1

		remaining := seed[1]
		start := seed[0]
		for {
			if remaining <= 0 {
				break
			}

			startLocation, consumed := walk(start, remaining, 0, maps)
			remaining -= consumed
			start += consumed
			if lowest == -1 || startLocation < lowest {
				lowest = startLocation
			}
		}

		location <- lowest
	}
}

func walk(value int, rng int, mapIndex int, maps [7][]mapping) (int, int) {
	if mapIndex >= len(maps) {
		return value, rng
	}

	mapp := maps[mapIndex]
	var rangeItem *mapping = nil
	for _, item := range mapp {
		if value >= item.SourceStart && value < item.SourceStart+item.Length {
			rangeItem = &item
			break
		}
	}

	if rangeItem != nil {
		diff := value - rangeItem.SourceStart
		newValue := rangeItem.DestinationStart + diff
		return walk(newValue, int(math.Min(float64(rng), float64(rangeItem.Length-diff))), mapIndex+1, maps)
	}

	return walk(value, 1, mapIndex+1, maps)
}
