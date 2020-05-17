package main

import "fmt"

func next(n int, i int, dir int) (int, int) {

	if n == 1 {
		dir = dir
		return i, dir
	}
	if i == n-1 && dir == 1 {
		dir = -dir
	}
	if i == 0 && dir == -1 {
		dir = -dir
	}

	i += dir
	return i, dir
}
func reverse(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
func add(a []int, n int) []int {
	for i := range a {
		a[i] |= (1 << n)
	}
	return a
}
func grayCode(n int) []int {
	ret := make([]int, 0)
	if n == 0 {
		ret = append(ret, 0)
		return ret
	}
	if n == 1 {
		ret = append(ret, 0, 1)
		return ret
	}
	ret = grayCode(n - 1)
	return append(ret, add(reverse(ret), n-1)...)
}
func main() {
	fmt.Printf("%v\n", grayCode(3))
}
