package util

import "fmt"

func PrintMatrix(matrix [][]byte) {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				matrix[i][j] = ' '
			}

			fmt.Print(string(matrix[i][j]))
		}
		fmt.Println()
	}

	fmt.Println()
}

func PrintBoolMatrix(matrix [][]bool) {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}
