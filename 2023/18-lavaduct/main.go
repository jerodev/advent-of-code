package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {
	file := util.FileFromArgs()
	b := make([]byte, 1)
	line := ""

	var y, x int64 = 1, 1
	var perimeter int64
	points := [][2]int64{
		{1, 1},
	}

	// Part 2 variables
	var y2, x2 int64 = 1, 1
	var perimeter2 int64
	points2 := [][2]int64{
		{1, 1},
	}

	for {
		_, err := file.Read(b)
		if b[0] == '\n' || err == io.EOF {
			y, x, perimeter = dig(line, y, x, perimeter, false)
			points = append(points, [2]int64{y, x})

			y2, x2, perimeter2 = dig(line, y2, x2, perimeter2, true)
			points2 = append(points2, [2]int64{y2, x2})

			line = ""

			if err == io.EOF {
				break
			}

			continue
		}

		line += string(b[0])
	}

	fmt.Println(shoelaceSurface(points, perimeter))
	fmt.Println(shoelaceSurface(points2, perimeter2))
}

func dig(instructions string, y, x, perimeter int64, hex bool) (int64, int64, int64) {
	parts := strings.SplitN(instructions, " ", 3)
	direction := parts[0]
	d, _ := strconv.Atoi(parts[1])
	distance := int64(d)

	if hex {
		color := strings.Trim(parts[2], "(#)")
		switch color[len(color)-1] {
		case '0':
			direction = "R"
		case '1':
			direction = "D"
		case '2':
			direction = "L"
		case '3':
			direction = "U"
		}

		distance, _ = strconv.ParseInt(color[:len(color)-1], 16, 64)
	}

	perimeter += distance

	switch direction {
	case "U":
		y -= distance
	case "D":
		y += distance
	case "R":
		x += distance
	case "L":
		x -= distance
	}

	return y, x, perimeter
}

func shoelaceSurface(points [][2]int64, perimeter int64) int64 {
	var sum int64

	for i := 0; i < len(points)-1; i++ {
		sum += points[i][0]*points[i+1][1] - points[i][1]*points[i+1][0]
	}

	if sum < 0 {
		sum = -sum
	}

	return (sum+perimeter)/2 + 1
}
