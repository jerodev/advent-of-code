package main

import (
	"fmt"
	"math"
)

const (
	PlayerHitPoints = 50
	playerMana      = 500

	bossHitPoints = 58
	bossDamage    = 9
)

var spells = []spell{
	{cost: 53, damage: 4},            // Magic Missile
	{cost: 73, damage: 2, heal: 2},   // Drain
	{cost: 113, turns: 6, armor: 7},  // Shield
	{cost: 173, damage: 3, turns: 6}, // Poison
	{cost: 229, turns: 5, mana: 101}, // Recharge
}

func main() {
	var sl spellList
	m, maxMana := 0, math.MaxInt

	for len(sl.counter) < 1000 {
		sl.Next()
		m = battle(sl.counter)
		if m > 0 && m < maxMana {
			maxMana = m
			fmt.Println(sl.counter)
			fmt.Println(m)
		}
	}
}

type spell struct {
	cost,
	damage,
	turns, armor, heal, mana int
}

type spellList struct {
	counter []int
}

func (sl *spellList) Next() {
	if len(sl.counter) == 0 {
		sl.counter = []int{0}
		return
	}

	sl.counter[0]++

	if sl.counter[0] == len(spells) {
		sl.counter[0] = 0
		if len(sl.counter) == 1 {
			sl.counter = append(sl.counter, 0)
			return
		}

		for i := 1; i < len(sl.counter); i++ {
			sl.counter[i]++
			if sl.counter[i] < len(spells) {
				break
			}

			sl.counter[i] = 0
			if i == len(sl.counter)-1 {
				sl.counter = append(sl.counter, 0)
				break
			}
		}
	}
}

// battle simulates the battle and returns the amount of mana spent
// If the battle is lost, the result is zero.
func battle(sl []int) int {
	var i, e, t, armor, cost int
	var gotoProtect int
	effects := map[int]int{}
	var s spell
	var ok bool

	player, mana, boss := PlayerHitPoints, playerMana, bossHitPoints
	for {
		armor = 0

		// Player turn
		// Ongoing effects
		for e, t = range effects {
			if t == 0 {
				continue
			}

			if spells[e].damage > 0 {
				boss -= spells[e].damage
			} else if spells[e].heal > 0 {
				player += spells[e].heal
			} else if spells[e].armor > 0 {
				armor += spells[e].armor
			} else if spells[e].mana > 0 {
				mana += spells[e].mana
			}

			effects[e]--
		}
		if boss < 1 {
			return cost
		}

		// New spell
		gotoProtect = 0
	selectSpell:
		s = spells[sl[i%len(sl)]]
		if s.turns > 0 {
			t, ok = effects[sl[i%len(sl)]]
			if ok && t > 0 {
				i++
				gotoProtect++
				if gotoProtect > 20 {
					return 0
				}
				goto selectSpell
			}

			cost += s.cost
			mana -= s.cost
			effects[sl[i%len(sl)]] = s.turns
		} else {
			cost += s.cost
			mana -= s.cost

			boss -= s.damage
			if boss < 1 && mana >= 0 {
				return cost
			}
		}
		if mana < 0 {
			return 0
		}

		// Boss turn
		player -= bossDamage - armor
		if player < 1 {
			return 0
		}

		i++
	}
}
