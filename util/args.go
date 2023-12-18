package util

import (
	"io"
	"os"
)

// FileFromArgs returns a file reader based on the first command line argument
func FileFromArgs() *os.File {
	file := os.Args[1]

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	return f
}

func GridFromArgsInt() [][]int {
	file := FileFromArgs()

	grid := [][]int{
		{},
	}
	rowNumber := 0

	b := make([]byte, 1)

	for {
		_, err := file.Read(b)
		if err == io.EOF {
			break
		}

		if b[0] == '\n' {
			grid = append(grid, []int{})
			rowNumber++
			continue
		}

		grid[rowNumber] = append(grid[rowNumber], int(b[0])-0x30)
	}

	return grid
}

// ReadFileFromArgs returns all file content as a string
func ReadFileFromArgs() string {
	file := os.Args[1]

	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(f)
}
