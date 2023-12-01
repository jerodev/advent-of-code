package util

import (
	"io"
	"os"
)

// FileFromArgs returns a file reader based on the first command line argument
func FileFromArgs() io.ReadCloser {
	file := os.Args[1]

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	return f
}
