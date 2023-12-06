package util

import (
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

// ReadFileFromArgs returns all file content as a string
func ReadFileFromArgs() string {
	file := os.Args[1]

	f, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(f)
}
