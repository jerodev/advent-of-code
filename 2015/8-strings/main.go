package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

var regEscape = regexp.MustCompile(`\\(?:(x[\dA-Fa-f]{2})|.)`)
var regEncodeHex = regexp.MustCompile(`(\\x[\dA-Fa-f]{2})`)

func main() {
	lines := strings.Split(input, "\n")

	var diff, diff2 int

	for _, l := range lines {
		diff += len(l) - len(escape(l[1:len(l)-1]))
		diff2 += 2 + strings.Count(l, `"`) + strings.Count(l, `\`)
	}

	fmt.Println(diff)
	fmt.Println(diff2)
}

func escape(s string) string {
	return regEscape.ReplaceAllString(s, "0")
}
