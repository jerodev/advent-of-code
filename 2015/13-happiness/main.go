package main

import (
	"advent-of-code/util"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

var h = map[string]map[string]int{}
var people []string

func main() {
	// Read the input
	parts := strings.Split(input, "\n")
	var l []string
	var tmp int
	var ok bool
	for i := range parts {
		l = strings.Split(parts[i][:len(parts[i])-1], " ")

		if !slices.Contains(people, l[0]) {
			people = append(people, l[0])
		}

		tmp = util.MustAtoi(l[3])
		if l[2] == "lose" {
			tmp *= -1
		}

		if _, ok = h[l[0]]; ok {
			h[l[0]][l[10]] = tmp
		} else {
			h[l[0]] = map[string]int{l[10]: tmp}
		}
	}

	// Part 2: Add "myself"
	h["myself"] = map[string]int{}
	for _, p := range people {
		h["myself"][p] = 0
		h[p]["myself"] = 0
	}
	people = append(people, "myself")

	// Run all possible scenarios!
	var maxHappiness int
	for i := range people {
		tmp = seatPeople([]string{people[i]}, 0)
		if tmp > maxHappiness {
			maxHappiness = tmp
		}
	}

	fmt.Println(maxHappiness)
}

func seatPeople(table []string, happiness int) int {
	if len(table) == len(people) {
		// Our table is full, make sure we count the happiness for the first and last person
		return happiness + h[table[0]][table[len(table)-1]] + h[table[len(table)-1]][table[0]]
	}

	// Loop over all people this person can sit next to
	tmp, happy := 0, happiness
	for p, ha := range h[table[len(table)-1]] {
		if slices.Contains(table, p) {
			// Already seated
			continue
		}

		tmp = seatPeople(append(table, p), happiness+ha+h[p][table[len(table)-1]])
		if tmp > happy {
			happy = tmp
		}
	}

	return happy
}
