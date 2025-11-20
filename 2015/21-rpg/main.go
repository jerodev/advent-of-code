package main

import (
	"fmt"
	"math"
)

const (
	bossHp     = 104
	bossDamage = 8
	bossArmor  = 1
)

var shop [][]item = [][]item{
	{ // 1
		{8, 4, 0},  // Dagger
		{10, 5, 0}, // Shortsword
		{25, 6, 0}, // Warhammer
		{40, 7, 0}, // Longsword
		{74, 8, 0}, // Greataxe
	},
	{ // 0..1
		{13, 0, 1},  // Leather
		{31, 0, 2},  // Chainmail
		{53, 0, 3},  // Splintmail
		{75, 0, 4},  // Bandedmail
		{102, 0, 5}, // Platemail
	},
	{ // 0..2
		{20, 0, 1},  // Defense +1
		{25, 1, 0},  // Damage +1
		{40, 0, 2},  // Defense +2
		{50, 2, 0},  // Damage +2
		{80, 0, 3},  // Defense +3
		{100, 3, 0}, // Damage +3
	},
}

func main() {
	minCoins, maxCoins := math.MaxInt, 0
	var dmg, def, cost int
	var admg, adef, acost int
	var rdmg, rdef, rcost, rings int

	ringRange := int(math.Pow(2, float64(len(shop[2]))))

	// Loop over all weapons, starting with the cheapest
	for w := range shop[0] {
		dmg, def, cost = shop[0][w].damage, shop[0][w].armor, shop[0][w].cost

		// Add a piece of armor
		for a := -1; a < len(shop[1]); a++ {
			if a > -1 {
				admg = dmg + shop[1][a].damage
				adef = def + shop[1][a].armor
				acost = cost + shop[1][a].cost
			} else {
				admg = dmg
				adef = def
				acost = cost
			}

			// Add rings
			for r := range ringRange {
				rdmg = admg
				rdef = adef
				rcost = acost
				rings = 0

				for ri := range shop[2] {
					if r>>ri&1 == 1 {
						rdmg += shop[2][ri].damage
						rdef += shop[2][ri].armor
						rcost += shop[2][ri].cost
						rings++
						if rings == 2 {
							break
						}
					}
				}

				if battle(rdmg, rdef) {
					minCoins = min(rcost, minCoins)
				} else {
					maxCoins = max(rcost, maxCoins)
				}
			}
		}
	}

	fmt.Println(minCoins)
	fmt.Println(maxCoins)
}

type item struct {
	cost, damage, armor int
}

// battle simulates the entire battle and indicates if the player has won
func battle(playerDamage, playerArmor int) bool {
	boss, player := bossHp, 100

	for {
		boss -= max(1, playerDamage-bossArmor)
		if boss < 1 {
			return true
		}

		player -= max(1, bossDamage-playerArmor)
		if player < 1 {
			return false
		}

		//fmt.Printf("Boss: %v\t\tPlayer: %v\n", boss, player)
	}
}
