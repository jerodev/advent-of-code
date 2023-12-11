package util

import "fmt"

func PrintMatrix(matrix [][]byte) {
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Print(string(matrix[i][j]))
		}
		fmt.Println()
	}
}
