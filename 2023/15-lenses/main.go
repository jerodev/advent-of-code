package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type lense struct {
	Label       string
	FocalLength int
}

type boxList map[int][]lense

func main() {
	file := util.FileFromArgs()

	b := make([]byte, 1)
	line := ""
	sum := 0

	boxes := boxList{}

	for {
		_, err := file.Read(b)
		if err == io.EOF {
			sum += hash(line)
			boxes = box(boxes, line)
			break
		}

		if b[0] == ',' {
			sum += hash(line)
			boxes = box(boxes, line)
			line = ""
			continue
		}

		line += string(b)
	}

	fmt.Println(sum)
	fmt.Println(calculatePower(boxes))
}

func hash(in string) int {
	current := 0

	for _, c := range in {
		current += int(c)
		current *= 17
		current %= 256
	}

	return current
}

func box(boxes boxList, op string) boxList {
	if strings.HasSuffix(op, "-") {
		return removeLens(boxes, op[:len(op)-1])
	}

	if strings.Contains(op, "=") {
		parts := strings.Split(op, "=")
		length, _ := strconv.Atoi(parts[1])

		return addLens(boxes, hash(parts[0]), lense{
			Label:       parts[0],
			FocalLength: length,
		})
	}

	return boxes
}

func removeLens(boxes boxList, label string) boxList {
	for i := range boxes {
		for j := range boxes[i] {
			if boxes[i][j].Label == label {
				boxes[i] = append(boxes[i][:j], boxes[i][j+1:]...)

				if len(boxes[i]) == 0 {
					delete(boxes, i)
				}

				return boxes
			}
		}
	}

	return boxes
}

func addLens(boxes boxList, boxNumber int, lens lense) boxList {
	if _, ok := boxes[boxNumber]; !ok {
		boxes[boxNumber] = []lense{lens}
		return boxes
	}

	for i := range boxes[boxNumber] {
		if boxes[boxNumber][i].Label == lens.Label {
			boxes[boxNumber][i].FocalLength = lens.FocalLength
			return boxes
		}
	}

	boxes[boxNumber] = append(boxes[boxNumber], lens)

	return boxes
}

func calculatePower(boxes boxList) int {
	sum := 0

	for i := range boxes {
		for j := range boxes[i] {
			sum += (i + 1) * (j + 1) * boxes[i][j].FocalLength
		}
	}

	return sum
}
