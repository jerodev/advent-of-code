package main

import (
	"advent-of-code/util"
	"fmt"
)

const empty = -1

func main() {
	input := util.StringToInts(util.ReadFileFromArgs(), "")

	var fragmented []int
	var id int

	for i := 0; i < len(input); i++ {
		// Add the next file
		for range input[i] {
			fragmented = append(fragmented, id)
		}

		id++

		// Add empty space
		i++
		if i < len(input) {
			for range input[i] {
				fragmented = append(fragmented, empty)
			}
		}
	}

	// Part 1: Give it a few rounds, to ensure no memory holes are left
	pt1 := make([]int, len(fragmented))
	copy(pt1, fragmented)
	fmt.Println(checksum(defrag(defrag(pt1))))

	// Part 2:
	fmt.Println(checksum(defragFiles(fragmented, id-1)))
}

func defrag(fragmented []int) []int {
	ptr, rptr := 0, len(fragmented)-1
	for {
		if fragmented[ptr] == empty {
			for ; ; rptr-- {
				if fragmented[rptr] != empty {
					break
				}
			}

			fragmented[ptr] = fragmented[rptr]
			fragmented[rptr] = empty
		}

		ptr++
		if ptr >= rptr {
			break
		}
	}

	return fragmented
}

func defragFiles(fragmented []int, fileId int) []int {
	fmt.Println(fragmented)

	rptr := len(fragmented) - 1
	var memStart, memSize, ptr int

	for ; fileId >= 0; fileId-- {
		// Find the memory block for this file
		for ; fragmented[rptr] != fileId; rptr-- {
		}
		memStart = rptr
		for ; rptr >= 0 && fragmented[rptr] == fileId; rptr-- {
		}
		memSize = memStart - rptr

		// Find a fitting block of free memory
		for ptr = 0; ptr <= rptr; ptr++ {
			for ; ptr <= rptr && fragmented[ptr] != empty; ptr++ {
			}
			memStart = ptr
			for ; ptr <= rptr && fragmented[ptr] == empty; ptr++ {
			}

			// Free memory found, fill it!
			if ptr-memStart >= memSize {
				for i := range memSize {
					fragmented[memStart+i] = fileId
					fragmented[rptr+memSize-i] = empty
				}

				break
			}
		}
	}

	fmt.Println(fragmented)

	return fragmented
}

func checksum(memory []int) int {
	var checksum int
	for i := range memory {
		if memory[i] == empty {
			continue
		}

		checksum += i * memory[i]
	}

	return checksum
}
