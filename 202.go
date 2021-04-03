package main

import (
	"fmt"
	"strconv"
)

func isHappy(n int) bool {
	str := strconv.Itoa(n)
	m := make(map[int]bool)
	m[n] = true
	for {
		if n == 1 {
			return true
		}
		s := 0
		for _, c := range []byte(str) {
			// calc n
			a := int(c - '0')
			s += a * a
		}
		n = s
		if m[n] && n != 1 {
			return false
		}
		m[n] = true
	}
}
func main() {
	fmt.Printf("%v\n", isHappy(19))
}
