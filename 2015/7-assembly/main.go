package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// Override b
// Set 0 for part 1 result
const override uint16 = 3176

func main() {
	reg := map[string]uint16{}
	var parts []string
	var ok, ok2 bool
	var tmp, tmp2 uint16
	var i string

	inst := queue[string]{
		q: strings.Split(input, "\n"),
	}

	for inst.HasNext() {
		i = inst.Shift()
		parts = strings.Split(i, " ")

		// Direct assignment
		// 123 -> x
		if len(parts) == 3 {
			tmp, ok = resolveInput(reg, parts[0])
			if ok {
				reg[parts[2]] = tmp
			} else {
				inst.Push(i)
			}

			continue
		}

		// NOT e -> f
		if parts[0] == "NOT" {
			tmp, ok = resolveInput(reg, parts[1])
			if ok {
				reg[parts[3]] = ^tmp
			} else {
				inst.Push(i)
			}

			continue
		}

		// a AND b -> c
		tmp, ok = resolveInput(reg, parts[0])
		tmp2, ok2 = resolveInput(reg, parts[2])
		if !ok || !ok2 {
			inst.Push(i)
			continue
		}
		switch parts[1] {
		case "AND":
			reg[parts[4]] = tmp & tmp2
		case "LSHIFT":
			reg[parts[4]] = tmp << tmp2
		case "RSHIFT":
			reg[parts[4]] = tmp >> tmp2
		case "OR":
			reg[parts[4]] = tmp | tmp2
		}
	}

	for x := range reg {
		fmt.Printf("%s:\t%v\n", x, reg[x])
	}

	fmt.Println(reg["a"])
}

func resolveInput(reg map[string]uint16, key string) (uint16, bool) {
	if override > 0 && key == "b" {
		return override, true
	}

	i, err := strconv.Atoi(key)
	if err == nil {
		return uint16(i), true
	}

	u16, ok := reg[key]
	return u16, ok
}

type queue[T any] struct {
	q []T
}

func (q *queue[T]) Push(v T) {
	q.q = append(q.q, v)
}

func (q *queue[T]) Shift() T {
	v := q.q[0]
	q.q = q.q[1:]

	return v
}

func (q *queue[T]) HasNext() bool {
	return len(q.q) > 0
}
