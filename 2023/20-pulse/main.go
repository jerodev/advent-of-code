package main

import (
	"advent-of-code/util"
	"fmt"
	"io"
	"strings"
)

const (
	MODULE_FLIP_FLOP = '%'
	MODULE_REMEMBER  = '&'
)

const (
	PULSE_LOW  = false
	PULSE_HIGH = true
)

type module struct {
	name         string
	moduleType   byte
	destinations []string
	memory       bool
}

func (m *module) propagate(p pulse) ([]string, bool) {
	if m.moduleType == MODULE_FLIP_FLOP {
		if p.level == PULSE_LOW {
			m.memory = !m.memory

			return m.destinations, m.memory
		}

		return []string{}, m.memory
	}

	if m.moduleType == MODULE_REMEMBER {
		m.memory = m.memory && p.level
		return m.destinations, !m.memory
	}

	// Broadcast
	return m.destinations, p.level
}

type pulse struct {
	level       bool
	destination *module
}

func main() {
	file := util.FileFromArgs()
	b := make([]byte, 1)

	modules := map[string]*module{}
	line := ""

	for {
		_, err := file.Read(b)
		if b[0] == '\n' || err == io.EOF {
			parts := strings.SplitN(line, " -> ", 2)
			name := parts[0]
			if name != "broadcaster" {
				name = parts[0][1:]
			}

			modules[name] = &module{
				name:         name,
				moduleType:   parts[0][0],
				destinations: strings.Split(parts[1], ", "),
				memory:       parts[0][0] == MODULE_REMEMBER,
			}

			if err == io.EOF {
				break
			}

			line = ""
			continue
		}

		line += string(b[0])
	}

	var high, low int64 = 0, 0
	var pulses []pulse
	var p pulse
	for b := 0; b < 4; b++ {
		pulses = []pulse{
			{
				level:       PULSE_LOW,
				destination: modules["broadcaster"],
			},
		}
		low++

		for {
			if len(pulses) == 0 {
				break
			}

			p = pulses[0]
			pulses = pulses[1:]

			destinations, hilo := p.destination.propagate(p)
			if hilo == PULSE_HIGH {
				high += int64(len(destinations))
			} else {
				low += int64(len(destinations))
			}

			for _, d := range destinations {
				fmt.Printf("%v -%v-> %v\n", p.destination.name, hilo, d)
				if d, ok := modules[d]; ok {
					pulses = append(pulses, pulse{
						level:       hilo,
						destination: d,
					})
				}
			}
		}

		// Reset memory modules
		for _, m := range modules {
			if m.moduleType == MODULE_REMEMBER {
				m.memory = true
			}
		}

		fmt.Println()
	}

	fmt.Println(high, low, high*low)
}
