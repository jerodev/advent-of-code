package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

const input = "yzbqklnj"

func main() {
	var i, fz int
	var s []byte

	var m5 [16]byte
	for {
		s = []byte(input + strconv.Itoa(i))
		m5 = md5.Sum(s)

		if m5[0] == 0 && m5[1] == 0 && m5[2] < 16 {
			if fz == 0 {
				fz = i
			}

			if m5[2] == 0 {
				break
			}
		}

		i++
	}

	fmt.Println(fz)
	fmt.Println(i)
}
