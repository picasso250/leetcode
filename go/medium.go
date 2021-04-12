package main

import (
	"strconv"
	"strings"
)

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
			if newItv, isIt := isInter(first, itv); isIt {
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
func isInter(a, b []int) (newInterval []int, isInter bool) {
	isInter = isBetween(a[0], b) || isBetween(a[1], b) ||
		isBetween(b[0], a) || isBetween(b[1], a)
	newInterval = []int{minInts(a[0], b[0]), maxInts(a[1], b[1])}
	return
}
func isBetween(a int, interval []int) bool {
	return interval[0] <= a && a <= interval[1]
}

// func insert(intervals [][]int, newInterval []int) [][]int {
// 	i, j := 0, len(intervals)
// 	for i < j {
// 		h := (i + j) / 2
// 		if newInterval[0] <= intervals[h][0] {
// 			i = h + 1
// 		} else {
// 			j = h
// 		}
// 	}
// 	j := -1
// 	k := -1
// 	for {
// 		if newItv, isIt := isInter(itv, newInterval); isIt {
// 			if j == -1 {
// 				j = i
// 			}
// 			newInterval = newItv
// 		} else {
// 			if j >= 0 && k == -1 {
// 				k = i
// 			}
// 			if newInterval[0] > itv[0] {
// 				j = i + 1
// 				k = i + 1
// 			}
// 		}
// 	}
// 	if j < 0 {
// 		j, k = 0, 0
// 	}
// 	if k == j {
// 		intervals = append(intervals, []int{})
// 		copy(intervals[j+1:], intervals[k:len(intervals)-1])
// 	} else {
// 		copy(intervals[j+1:], intervals[k:])
// 	}
// 	intervals[j] = newInterval
// 	return intervals[:j+len(intervals)-k]
// }
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return numDecodings1(s[0])
	}
	// if len(s) == 2 {
	// 	return numDecodings2(s[0], s[1])
	// }
	if s[0] == '0' {
		return 0
	}
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	dp[n-1] = numDecodings1(s[n-1])
	// dp[n-2] = numDecodings2(s[n-2], s[n-1])
	for i := n - 2; i >= 0; i-- {
		cnt := 0
		if s[i] != '0' {
			cnt += dp[i+1]
		} else {
			dp[i] = 0
			continue
		}
		a, _ := strconv.Atoi(s[i : i+2])
		if a > 0 && a <= 26 {
			cnt += dp[i+2]
		}
		// if !(s[i] != '0') && !(a > 0 && a <= 26) {
		// 	return 0
		// }
		dp[i] = cnt
	}
	return dp[0]
}
func numDecodings1(b byte) int {
	if b != '0' {
		return 1
	}
	return 0
}

// func numDecodings2(a, b byte) int {
// 	n, _ := strconv.Atoi(string([]byte{a, b}))
// 	cnt := 0
// 	if a != '0' && b != '0' {
// 		cnt++
// 	}
// 	if n > 0 && n <= 26 {
// 		cnt++
// 	}
// 	return cnt
// }
