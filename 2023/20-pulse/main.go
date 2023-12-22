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
	moduleType   byte
	destinations []string
	memory       bool
}

func (m *module) propagate(p pulse) ([]string, bool) {
	if m.moduleType == MODULE_FLIP_FLOP {
		if p.level == PULSE_LOW {
			if m.memory == PULSE_HIGH {
				m.memory = PULSE_LOW
			} else {
				m.memory = PULSE_HIGH
			}
		}

		return m.destinations, m.memory
	}

	// Remember module
	m.memory = m.memory && p.level
	if m.memory {
		return m.destinations, PULSE_LOW
	}

	return m.destinations, PULSE_HIGH
}

type pulse struct {
	level       bool
	destination *module
}

type pulseQueue []pulse

func (q *pulseQueue) pop() pulse {
	p := (*q)[0]
	*q = (*q)[1:]

	return p
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
				moduleType:   parts[0][0],
				destinations: strings.Split(parts[1], ", "),
				memory:       true,
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
	for b := 0; b < 1000; b++ {
		pulses := pulseQueue{
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

			p := pulses.pop()

			destinations, hilo := p.destination.propagate(p)
			if hilo == PULSE_HIGH {
				high += int64(len(destinations))
			} else {
				low += int64(len(destinations))
			}

			for _, d := range destinations {
				if d, ok := modules[d]; ok {
					pulses = append(pulses, pulse{
						level:       hilo,
						destination: d,
					})
				}
			}
		}
	}

	fmt.Println(high, low, high*low)
}
