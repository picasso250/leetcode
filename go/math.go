package main

func minInts(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, x := range a {
		if x < min {
			min = x
		}
	}
	return min
}
func maxInts(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	max := a[0]
	for _, x := range a {
		if x > max {
			max = x
		}
	}
	return max
}
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
