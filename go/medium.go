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
