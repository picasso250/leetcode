package main

import "strings"

func simplifyPath(path string) string {
	a := strings.Split(path, "/")
	b := []string{a[0]}
	for i := 1; i < len(a); i++ {
		f := a[i]
		if f == "" {
			// do nothing
		} else if f == "." {
			// do nothing
		} else if f == ".." {
			if len(b) > 1 {
				b = b[:len(b)-1]
			} // else do nothing
		} else {
			b = append(b, f)
		}
	}
	if len(b) == 1 {
		return "/"
	}
	return strings.Join(b, "/")
}
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	i, j := 0, m*n
	for i < j {
		h := (i + j) / 2
		if matrix[h/n][h%n] == target {
			return true
		} else if matrix[h/n][h%n] < target {
			i = h + 1
		} else {
			j = h
		}
	}
	return false
}
func merge(intervals [][]int) [][]int {
	i := 0
	n := len(intervals)
outer:
	for i < n {
		first := intervals[i]
		rest := intervals[i+1 : n]
		for j, itv := range rest {
			if newItv, isIt := mergeInter(first, itv); isIt {
				intervals[i] = newItv
				if len(rest) > 0 {
					rest[j] = rest[len(rest)-1]
				}
				// rest = rest[:len(rest)-1]
				n--
				continue outer
			}
		}
		i++
	}
	return intervals[:i]
}
func mergeInter(a, b []int) (newInterval []int, isInter bool) {
	isInter = mergeBetween(a[0], b) || mergeBetween(a[1], b) ||
		mergeBetween(b[0], a) || mergeBetween(b[1], a)
	newInterval = []int{minInts(a[0], b[0]), maxInts(a[1], b[1])}
	return
}
func mergeBetween(a int, interval []int) bool {
	return interval[0] <= a && a <= interval[1]
}
