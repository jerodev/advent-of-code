package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type part struct {
	X, M, A, S int
}

func (p part) Sum() int64 {
	return int64(p.X + p.M + p.A + p.S)
}

type rule struct {
	Target       byte
	Function     byte
	Number       int
	NextWorkflow string
}

func (r rule) eval(p part) bool {
	var value int
	switch r.Target {
	case 'x':
		value = p.X
	case 'm':
		value = p.M
	case 'a':
		value = p.A
	case 's':
		value = p.S
	default:
		return true
	}

	switch r.Function {
	case '<':
		return value < r.Number
	case '>':
		return value > r.Number
	}

	return false
}

type workflow struct {
	Rules []rule
}

func main() {
	file := util.FileFromArgs()

	workflows := map[string]workflow{}
	b := make([]byte, 1)
	line := ""

	// Parse workflows
	for {
		_, _ = file.Read(b)
		if b[0] == '\n' {
			if line == "" {
				break
			}

			parts := strings.SplitN(line, "{", 2)
			ruleParts := strings.Split(parts[1][:len(parts[1])-1], ",")
			var rules []rule
			sameNext := true
			lastNext := ""
			for _, rulePart := range ruleParts {
				if !strings.Contains(rulePart, ":") {
					rules = append(rules, rule{
						Target:       '_',
						NextWorkflow: rulePart,
					})
				} else {
					pts := strings.Split(rulePart, ":")
					number, _ := strconv.Atoi(pts[0][2:])
					rulePart = pts[1]
					rules = append(rules, rule{
						Target:       pts[0][0],
						Function:     pts[0][1],
						Number:       number,
						NextWorkflow: rulePart,
					})
				}

				if sameNext && lastNext != "" && lastNext != rulePart {
					sameNext = false
				}
				lastNext = rulePart
			}

			// If all conditions lead to the same workflow, skip checking for rules
			if sameNext {
				workflows[parts[0]] = workflow{
					Rules: rules[len(rules)-1:],
				}
			} else {
				workflows[parts[0]] = workflow{
					Rules: rules,
				}
			}

			line = ""

			continue
		}

		line += string(b[0])
	}

	// Parse parts
	var parts []part
	for {

		_, err := file.Read(b)
		if b[0] == '\n' || err == io.EOF {
			line = strings.Trim(line, "{}\n")
			pts := strings.SplitN(line, ",", 4)

			parts = append(parts, part{
				X: util.MustAtoi(pts[0][2:]),
				M: util.MustAtoi(pts[1][2:]),
				A: util.MustAtoi(pts[2][2:]),
				S: util.MustAtoi(pts[3][2:]),
			})

			line = ""

			if err == io.EOF {
				break
			}
		}

		line += string(b[0])
	}

	// Part 1
	var sum int64 = 0
	for _, part := range parts {
		workflowKey := "in"
	outer:
		for {
			workflow := workflows[workflowKey]
			for _, rule := range workflow.Rules {
				if rule.eval(part) {
					if rule.NextWorkflow == "A" {
						sum += part.Sum()
						break outer
					}
					if rule.NextWorkflow == "R" {
						break outer
					}

					workflowKey = rule.NextWorkflow
					break
				}
			}
		}
	}

	fmt.Println(sum)

	// Part 2
	sum = 0
	for x := 1; x <= 4000; x++ {
		for m := 1; m <= 4000; m++ {
			for a := 1; a <= 4000; a++ {
				for s := 1; s <= 4000; s++ {
					workflowKey := "in"
				outer2:
					for {
						workflow := workflows[workflowKey]
						for _, rule := range workflow.Rules {
							if rule.eval(part{X: x, M: m, A: a, S: s}) {
								if rule.NextWorkflow == "A" {
									sum += int64(x + m + a + s)
									break outer2
								}
								if rule.NextWorkflow == "R" {
									break outer2
								}

								workflowKey = rule.NextWorkflow
								break
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(sum)
}
