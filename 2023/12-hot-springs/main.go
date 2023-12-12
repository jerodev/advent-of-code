package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"strings"
)

const (
	DAMAGE      = '#'
	OPERATIONAL = '.'
	UNKNOWN     = '?'
)

func main() {
	fmt.Println(run(util.FileFromArgs(), 1))
	fmt.Println(run(util.FileFromArgs(), 5))
}

var fitNumberCache = map[string]int{}

func run(file io.Reader, foldTimes int) int {
	b := make([]byte, 1)
	line := ""

	result := 0
	lineNumber := 0

	for {
		_, err := file.Read(b)
		if b[0] == "\n"[0] || err == io.EOF {
			parts := strings.SplitN(line, " ", 2)
			numbers := util.StringToInts(parts[1], ",")

			history := strings.Repeat(strings.TrimSpace(parts[0])+"?", foldTimes)
			history = history[:len(history)-1]

			allNumbers := []int{}
			for i := 0; i < foldTimes; i++ {
				allNumbers = append(allNumbers, numbers...)
			}

			count := calculateArrangements(history, allNumbers)
			lineNumber++
			fmt.Printf("%v - %v possibilities\n\n", lineNumber, count)
			result += count

			if err == io.EOF {
				break
			}

			line = ""
		}

		line += string(b)
	}

	fmt.Println()

	return result
}

func calculateArrangements(line string, numbers []int) int {
	groups, numbers := splitInGroups(line, numbers)
	fmt.Println(groups, numbers)

	return fitNumbers(groups, numbers)
}

func splitInGroups(line string, numbers []int) ([]string, []int) {
	groups := []string{}
	part := ""
	for _, b := range line {
		if byte(b) == OPERATIONAL {
			if len(part) > 0 {
				groups = append(groups, part)
				part = ""
			}

			continue
		}

		part += string(b)
	}

	if len(part) > 0 {
		groups = append(groups, part)
	}

	// If there are as many groups with damage as there are numbers, we can remove the groups without damage
	damageCounter := 0
	for _, g := range groups {
		if strings.IndexByte(g, DAMAGE) >= 0 {
			damageCounter++
		}
	}
	if damageCounter == len(numbers) {
		newGroups := []string{}
		for _, g := range groups {
			if strings.IndexByte(g, DAMAGE) >= 0 {
				newGroups = append(newGroups, g)
			}
		}

		groups = newGroups
	}

	// We can actually sanitize the groups a bunch.
	for {
		if len(groups) == 0 || len(numbers) == 0 {
			break
		}

		// If the first group is shorter than the first number, remove the group
		if len(groups[0]) < numbers[0] {
			groups = groups[1:]
			continue
		}

		// If the first group has damage and matches the number, we can remove it.
		if strings.IndexByte(groups[0], DAMAGE) > -1 && len(groups[0]) == numbers[0] {
			groups = groups[1:]
			numbers = numbers[1:]
			continue
		}

		// If the first group starts with damage, we can remove the first damage group
		if groups[0][0] == DAMAGE {
			if len(groups[0]) == numbers[0] || len(groups[0]) == numbers[0]+1 {
				groups = groups[1:]
				numbers = numbers[1:]
			} else {
				groups[0] = groups[0][numbers[0]+1:]
				numbers = numbers[1:]
			}

			continue
		}

		// If the first group starts with an UNKNOWN followed by a DAMAGE, we can remove the first damage group if the size matches
		if len(groups[0]) > 1 && groups[0][0] == UNKNOWN && groups[0][1] == DAMAGE {
			damageLength := 0
			for i := 1; i < len(groups[0]); i++ {
				if groups[0][i] == DAMAGE {
					damageLength++
				} else {
					break
				}
			}

			if damageLength == numbers[0] {
				groupLength := len(groups[0])
				if groupLength == damageLength+1 || groupLength == damageLength+2 {
					groups = groups[1:]
					numbers = numbers[1:]
				} else {
					groups[0] = groups[0][damageLength+2:]
					numbers = numbers[1:]
				}

				continue
			}
		}

		break
	}
	for {
		if len(groups) == 0 || len(numbers) == 0 {
			break
		}

		lastGroup := groups[len(groups)-1]
		lastNumber := numbers[len(numbers)-1]

		// If the last group is shorter than the last number, remove the group
		if len(lastGroup) < lastNumber {
			groups = groups[:len(groups)-1]
			continue
		}

		// If the last group has damage and matches the number, we can remove it.
		if strings.IndexByte(lastGroup, DAMAGE) > -1 && len(lastGroup) == lastNumber {
			groups = groups[:len(groups)-1]
			numbers = numbers[:len(numbers)-1]
			continue
		}

		// If the last group ends with damage, we can remove the last number
		if lastGroup[len(lastGroup)-1] == DAMAGE {
			if len(lastGroup) == lastNumber || len(lastGroup) == lastNumber+1 {
				groups = groups[:len(groups)-1]
				numbers = numbers[:len(numbers)-1]
			} else {
				groups[len(groups)-1] = lastGroup[:len(lastGroup)-lastNumber-1]
				numbers = numbers[:len(numbers)-1]
			}
			continue
		}

		// If the last group ends with an UNKNOWN followed by a DAMAGE, we can remove the last damage group if the size matches
		if len(lastGroup) > 1 && lastGroup[len(lastGroup)-1] == UNKNOWN && lastGroup[len(lastGroup)-2] == DAMAGE {
			damageLength := 0
			for i := len(lastGroup) - 2; i >= 0; i-- {
				if lastGroup[i] == DAMAGE {
					damageLength++
				} else {
					break
				}
			}

			if damageLength == lastNumber {
				groupLength := len(lastGroup)
				if groupLength == damageLength+1 || groupLength == damageLength+2 {
					groups = groups[:len(groups)-1]
					numbers = numbers[:len(numbers)-1]
				} else {
					groups[len(groups)-1] = lastGroup[:len(lastGroup)-lastNumber-2]
					numbers = numbers[:len(numbers)-1]
				}

				continue
			}
		}

		break
	}

	return groups, numbers
}

// fitNumbers returns the times numbers can fit in the given groups
func fitNumbers(groups []string, numbers []int) int {
	if len(numbers) == 0 {
		// If a group still contains unused damage, we cannot fit the numbers
		for _, group := range groups {
			if strings.IndexByte(group, DAMAGE) >= 0 {
				return 0
			}
		}

		return 1
	}
	if len(groups) == 0 {
		return 0
	}

	// Number won't fit current group, continue to the next group
	if strings.IndexByte(groups[0], DAMAGE) == -1 && len(groups[0]) < numbers[0] {
		return fitNumbers(groups[1:], numbers)
	}

	// Find the score in cache
	cacheSlug := strings.Join(groups, "|") + "|" + strings.Join(util.IntsToStrings(numbers), ",")
	if results, ok := fitNumberCache[cacheSlug]; ok {
		return results
	}

	results := 0
	for i := 0; i <= len(groups[0])-numbers[0]; i++ {
		// Skipping damage is not allowed
		if i > 0 && groups[0][i-1] == DAMAGE {
			break
		}
		// Next character cannot be damage!
		if i+numbers[0] < len(groups[0]) && groups[0][i+numbers[0]] == DAMAGE {
			continue
		}

		// Are we at the end of the current group?
		if i+numbers[0] == len(groups[0]) || i+numbers[0] == len(groups[0])-1 {
			results += fitNumbers(groups[1:], numbers[1:])
		} else if len(groups) > 1 {
			results += fitNumbers(append([]string{groups[0][i+numbers[0]+1:]}, groups[1:]...), numbers[1:])
		} else {
			results += fitNumbers([]string{groups[0][i+numbers[0]+1:]}, numbers[1:])
		}
	}

	// If the current group does not contain damage, continue to the next group
	if strings.IndexByte(groups[0], DAMAGE) == -1 && len(groups) > 1 {
		results += fitNumbers(groups[1:], numbers)
	}

	// Cache the result
	fitNumberCache[cacheSlug] = results

	return results
}
