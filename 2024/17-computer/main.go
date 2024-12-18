package main

import (
	"advent-of-code/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var A, B, C, ogB, ogC int
var pc int
var instructions []int
var output []int

func main() {
	file := util.FileFromArgs()
	scan := bufio.NewScanner(file)

	// Read input
	scan.Scan()
	A, _ = strconv.Atoi(scan.Text()[12:])
	scan.Scan()
	ogB, _ = strconv.Atoi(scan.Text()[12:])
	scan.Scan()
	ogC, _ = strconv.Atoi(scan.Text()[12:])
	scan.Scan()
	scan.Scan()
	instructions = util.StringToInts(scan.Text()[9:], ",")

	fmt.Println(A, ogB, ogC, instructions)
	B, C = ogB, ogC

	// Part 1: Run the program
	for Tick() {
	}
	fmt.Println(strings.Join(util.IntsToStrings(output), ","))

	// Part 2: Find the correct value for A
part2:
	for i := 0; ; i++ {
		A, B, C, pc = i, ogB, ogC, 0
		output = []int{}

		for Tick() {
			if len(output) > len(instructions) {
				continue part2
			}
		}

		if len(output) == len(instructions) {
			for j := range output {
				if output[j] != instructions[j] {
					continue part2
				}
			}

			fmt.Println(i)
			break
		}

		if i%1e7 == 0 {
			fmt.Println(i / 1e7)
		}
	}
}

// Tick executes the next instruction. The function will return true as long as the program is active
func Tick() bool {
	if pc > len(instructions)-1 {
		return false
	}

	operand := instructions[pc+1]
	if instructions[pc] != 1 && instructions[pc] != 3 && instructions[pc] != 4 {
		switch operand {
		case 4:
			operand = A
		case 5:
			operand = B
		case 6:
			operand = C
		default:
		}
	}

	operations[instructions[pc]](operand)

	pc += 2

	return true
}

var operations = [8]func(int){
	0: func(operand int) { // adv
		A = A / util.IntPow(2, operand)
	},
	1: func(operand int) { // bxl
		B = B ^ operand
	},
	2: func(operand int) { // bst
		B = operand % 8
	},
	3: func(operand int) { // jnz
		if A != 0 {
			pc = operand - 2
		}
	},
	4: func(operand int) { // bxc
		B = B ^ C
	},
	5: func(operand int) { // out
		output = append(output, operand%8)
	},
	6: func(operand int) { // bdv
		B = A / util.IntPow(2, operand)
	},
	7: func(operand int) { // cdv
		C = A / util.IntPow(2, operand)
	},
}
