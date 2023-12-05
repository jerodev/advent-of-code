package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"math"
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

func findSeedLocationWorker(maps [7][]mapping, seeds <-chan []int, location chan<- int) {
	for seed := range seeds {
		lowest := -1

		for r := 0; r < seed[1]; r++ {
			x := seed[0] + r
			for i := 0; i < len(maps); i++ {
				for _, mapping := range maps[i] {
					if x >= mapping.SourceStart && x < mapping.SourceStart+mapping.Length {
						x = mapping.DestinationStart + (x - mapping.SourceStart)
						break
					}
				}
			}

			if lowest == -1 || x < lowest {
				lowest = x
			}
		}

		location <- lowest
	}
}
