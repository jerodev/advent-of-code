package main

import (
	"advent-of-code/util"
	_ "embed"
	"fmt"
	"iter"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var ingredients []ingredient
	lines := strings.Split(input, "\n")
	var l []string
	for _, line := range lines {
		l = strings.Split(line, " ")
		ingredients = append(ingredients, ingredient{
			l[0],
			util.MustAtoi(l[2]),
			util.MustAtoi(l[4]),
			util.MustAtoi(l[6]),
			util.MustAtoi(l[8]),
			util.MustAtoi(l[10]),
		})
	}

	var maxScore, maxWithCalories int
	var dCapacity, dDurability, dFlavor, dTexture, dCalories, tmp int
	for d := range divisions() {
		dCapacity, dDurability, dFlavor, dTexture, dCalories = 0, 0, 0, 0, 0

		for i := range ingredients {
			dCapacity += ingredients[i].capacity * d[i]
			dDurability += ingredients[i].durability * d[i]
			dFlavor += ingredients[i].flavor * d[i]
			dTexture += ingredients[i].texture * d[i]
			dCalories += ingredients[i].calories * d[i]
		}

		if dCapacity <= 0 || dDurability <= 0 || dFlavor <= 0 || dTexture <= 0 {
			continue
		}

		tmp = dCapacity * dDurability * dFlavor * dTexture
		if tmp > maxScore {
			maxScore = tmp
		}

		if dCalories == 500 {
			if tmp > maxWithCalories {
				maxWithCalories = tmp
			}
		}
	}

	fmt.Println(maxScore)
	fmt.Println(maxWithCalories)
}

type ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
}

func divisions() iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		for i := range 100 - 2 {
			for j := range 100 - i {
				for k := range 100 - j {
					if !yield([]int{i, j, k, 100 - i - j - k}) {
						return
					}
				}
			}
		}
	}
}
