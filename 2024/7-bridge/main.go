package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

var modifiers = []func(int, int) int{
	func(i1, i2 int) int {
		return i1 + i2
	},
	func(i1, i2 int) int {
		if i1 == 0 {
			i1 = 1
		}

		return i1 * i2
	},
	func(i1, i2 int) int {
		return i1*util.IntPow(10, util.IntLength(i2)) + i2
	},
}

func main() {
	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)

	var parts []string
	var result, result2, sum int
	var numbers []int

	for scan.Scan() {
		parts = strings.SplitN(scan.Text(), ": ", 2)
		sum, _ = strconv.Atoi(parts[0])
		numbers = util.StringToInts(parts[1], " ")

		if resolves(sum, numbers, 2) {
			result += sum
			result2 += sum
		} else if resolves(sum, numbers, 3) {
			result2 += sum
		}
	}

	fmt.Println(result)
	fmt.Println(result2)
}

func resolves(sum int, numbers []int, base int) bool {
	possibilities := util.IntPow(base, len(numbers)-1)
	for i := range possibilities {
		bin := ("0000000000" + big.NewInt(int64(i)).Text(base))
		bin = bin[len(bin)-len(numbers)+1:]

		total := numbers[0]
		for k := range len(bin) {
			total = modifiers[bin[k]-'0'](total, numbers[k+1])

			if total > sum {
				break
			}
		}

		if total == sum {
			return true
		}
	}

	return false
}
