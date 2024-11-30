package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"math"
	"regexp"
	"slices"
	"strings"
)

var splitRegex = regexp.MustCompile(`\s*,\s*`)

const (
	RANGE_MIN = 200000000000000
	RANGE_MAX = 400000000000000
)

type hailstone struct {
	X, Y, Z    int64
	VX, VY, VZ int64
	M          float64
	C          float64
}

type point struct {
	X, Y float64
}

func (h hailstone) CollidesWith(h2 hailstone) point {
	// Two parallel lines never collide
	if h.M == h2.M {
		return point{}
	}

	var X float64
	if math.IsInf(h.M, 1) {
		X = h.C
	} else if math.IsInf(h2.M, 1) {
		X = h2.C
	} else {
		X = (h2.C - h.C) / (h.M - h2.M)
	}
	Y := h.M*X + h.C

	return point{
		X: X,
		Y: Y,
	}
}

func newHailstone(X, Y, Z, VX, VY, VZ int64) hailstone {
	X2, Y2 := X+VX, Y+VY

	M := math.Inf(1)
	if X2 != X {
		M = float64(Y2-Y) / float64(X2-X)
	}

	return hailstone{
		X:  X,
		Y:  Y,
		Z:  Z,
		VX: VX,
		VY: VY,
		VZ: VZ,
		M:  M,
		C:  float64(Y) - float64(X)*M,
	}
}

func main() {
	file := util.FileFromArgs()

	var stones []hailstone

	b := make([]byte, 1)
	line := ""
	for {
		_, err := file.Read(b)
		if b[0] == '\n' || err == io.EOF {
			parts := strings.SplitN(line, " @ ", 2)
			coords := splitRegex.Split(parts[0], 3)
			velocity := splitRegex.Split(parts[1], 3)
			stones = append(stones, newHailstone(
				int64(util.MustAtoi(coords[0])),
				int64(util.MustAtoi(coords[1])),
				int64(util.MustAtoi(coords[2])),
				int64(util.MustAtoi(velocity[0])),
				int64(util.MustAtoi(velocity[1])),
				int64(util.MustAtoi(velocity[2])),
			))

			line = ""

			if err == io.EOF {
				break
			}

			continue
		}

		line += string(b[0])
	}

	fmt.Println(stones)

	collisionCount := 0
	var history []string
	for i := range stones {
		for j := i + 1; j < len(stones); j++ {
			point := stones[i].CollidesWith(stones[j])
			if point.X == 0 && point.Y == 0 {
				continue
			}

			// Point already exists
			key := fmt.Sprintf("%v,%v", point.X, point.Y)
			if slices.Contains(history, key) {
				continue
			}

			// Within range
			if point.X < RANGE_MIN || point.X > RANGE_MAX || point.Y < RANGE_MIN || point.Y > RANGE_MAX {
				continue
			}

			collisionCount++
			history = append(history, key)
		}
	}

	fmt.Println(collisionCount)
}
