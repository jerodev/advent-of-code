package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"strings"
)

type brick struct {
	X1, Y1, Z1 int
	X2, Y2, Z2 int
}

func (bt *brick) isSupporting(b *brick) bool {
	// Needs to be one level above the supporting brick
	if util.Max(bt.Z1, bt.Z2)+1 != util.Min(b.Z1, b.Z2) {
		return false
	}

	// The shapes should cross each other
	return util.Max(bt.X1, b.X1) <= util.Min(bt.X2, b.X2) && util.Max(bt.Y1, b.Y1) <= util.Min(bt.Y2, b.Y2)
}

type tower []*brick

// canMove checks if a given brick can move down in the tower
func (t tower) canMove(b *brick) bool {
	// We are on the bottom
	if b.Z1 == 1 || b.Z2 == 1 {
		return false
	}

	for _, bt := range t {
		if bt.isSupporting(b) {
			return false
		}
	}

	return true
}

// settle moves bricks in the tower until it is stable
func (t tower) settle() {
	for {
		isStable := true

		for _, b := range t {
			if t.canMove(b) {
				b.Z1--
				b.Z2--
				isStable = false
			}
		}

		if isStable {
			break
		}
	}
}

func main() {
	file := util.FileFromArgs()

	b := make([]byte, 1)
	line := ""
	var bricks tower

	for {
		_, err := file.Read(b)
		if b[0] == '\n' || err == io.EOF {
			coords := strings.SplitN(line, "~", 2)
			coords1 := util.StringToInts(coords[0], ",")
			coords2 := util.StringToInts(coords[1], ",")

			bricks = append(bricks, &brick{
				X1: coords1[0],
				Y1: coords1[1],
				Z1: coords1[2],
				X2: coords2[0],
				Y2: coords2[1],
				Z2: coords2[2],
			})

			line = ""

			if err == io.EOF {
				break
			}

			continue
		}

		line += string(b[0])
	}

	// Let the bricks fall until there is no more movement
	bricks.settle()
	printTowerSide(bricks)

	// Now attempt to remove a brick without the tower collapsing
	newTower := make(tower, len(bricks)-1)
	stableBricks := 0
	for i := range bricks {
		newTower = append(append(tower{}, bricks[:i]...), bricks[i+1:]...)
		stable := true

		for _, b := range newTower {
			if newTower.canMove(b) {
				stable = false
				break
			}
		}

		if stable {
			stableBricks++
		}
	}

	fmt.Println(stableBricks)
}

func printTowerSide(t tower) {
	maxX, maxZ := 0, 0
	for _, b := range t {
		maxX = util.Max(util.Max(maxX, b.X1), b.X2)
		maxZ = util.Max(util.Max(maxZ, b.Z1), b.Z2)
	}

	grid := make([][]byte, maxZ)
	for i := range grid {
		grid[i] = make([]byte, maxX+1)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for i, b := range t {
		x := util.Min(b.X1, b.X2)
		dx := util.Abs(b.X1 - b.X2)
		z := util.Min(b.Z1, b.Z2) - 1
		dz := util.Abs(b.Z1 - b.Z2)

		for iz := 0; iz <= dz; iz++ {
			for ix := 0; ix <= dx; ix++ {
				grid[z+iz][x+ix] = byte(0x41 + i)
			}
		}
	}

	util.PrintMatrix(grid)
}
