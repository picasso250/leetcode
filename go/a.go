package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func max2Ints(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min2Ints(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	if m == 0 || n == 0 {
		return 0
	}
	var dp [][]int
	// init memory
	dp = make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// init dp
	for i := 0; i < n; i++ {
		dp[0][i] = 0
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 0
	}
	// calc
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max2Ints(dp[i][j+1], dp[i+1][j])
			}
		}
	}
	return dp[m][n]
}

type substr struct {
	length int
	pos    int
	str    string
	m      map[byte]bool
}

func findLongestSubStringBefore(s string, index int) substr {
	m := make(map[byte]bool)
	bs := make([]byte, 0)
	var i int
	for i = index; i >= 0; i-- {
		if m[s[i]] {
			break
		}
		bs = append([]byte{s[i]}, bs...)
		m[s[i]] = true
	}
	return substr{len(bs), i + 1, string(bs), m}
}

// O(n^2)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 1
	begin := 0
	end := 1
	m := map[byte]bool{s[begin]: true}
	for end < len(s) && begin+max < len(s) {
		c := s[end]
		if m[c] {
			m[s[begin]] = false
			begin++
		} else {
			m[c] = true
			end++
		}
		if end-begin > max {
			max = end - begin
		}
	}
	return max
}
func divide2sortedArray(a, b []int, i int) (cond bool) {
	m, n := len(a), len(b)
	j := m + n - (m+n)/2 - i
	// la, lb, ra, rb := []int{}, []int{}, []int{}, []int{}
	// if i > 0 {
	// 	la = a[:i]
	// }
	// if j > 0 {
	// 	lb = b[:j]
	// }
	// if m > i {
	// 	ra = a[i:m]
	// }
	// if n > j {
	// 	rb = b[j:n]
	// }
	// fmt.Println([][]int{la, ra})
	// fmt.Println([][]int{lb, rb})
	cond = abcond(a, b, i, j)
	// if len(lb) > 0 && len(ra) > 0 {
	// 	cond = append(cond, lb[len(lb)-1] <= ra[0])
	// }
	return cond
}
func abcond(a, b []int, i, j int) bool {
	// A[i-1]<=B[j]
	var aa = arrayGet(a, i-1)
	var bb = arrayGet(b, j)
	return aa <= bb
}
func arrayGet(a []int, i int) int {
	n := len(a)
	var aa int
	if i < 0 {
		aa = math.MinInt32
	} else if i >= 0 && i < n {
		aa = a[i]
	} else if i >= n {
		aa = math.MaxInt32
	}
	return aa
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m == 0 {
		if (m+n)%2 == 0 {
			return float64(nums2[n/2-1]+nums2[n/2]) / 2
		} else {
			return float64(nums2[n/2])
		}
	} else if n == 0 {
		if (m+n)%2 == 0 {
			return float64(nums1[m/2-1]+nums1[m/2]) / 2
		} else {
			return float64(nums1[m/2])
		}
	}
	// A[0]..A[i-1] | A[i]..A[m-1] ,  i=0 to m
	// B[0]..B[j-1] | B[j]..B[n-1]
	// 偶 (m+n)/2 | (m+n)/2
	// 奇 [(m+n)/2]+1 | [(m+n)/2]
	// m-1-(i-1) + n-1-(j-1) = [(m+n)/2]
	// m+n-(i+j)=[(m+n)/2]
	// A[i-1]<=B[j]
	// B[j-1]<=A[i]
	if m > n {
		m, n = n, m
		nums1, nums2 = nums2, nums1
	}
	// vv, _ := make([]bool, 0), make([]bool, 0)
	// i := 0
	// for i = 0; i <= m; i++ {
	// 	fmt.Println("--- i = " + strconv.Itoa(i) + " ---")
	// 	cond := divide2sortedArray(nums1, nums2, i)
	// 	fmt.Println(cond)
	// 	vv = append(vv, cond)
	// }
	// fmt.Println(vv)
	// for i = 0; i <= m; {
	// 	if vv[i] == true {
	// 		i++
	// 	} else {
	// 		break
	// 	}
	// }
	// i = i - 1
	// fmt.Println(i)
	// [true,false] find last true
	// trans to
	// [false,true]
	// find last false
	// Search find first true
	ii := sort.Search(m+1, func(ii int) bool {
		return divide2sortedArray(nums1, nums2, m-ii)
	})
	i := m - ii
	j := m + n - (m+n)/2 - i
	// conds := divide2sortedArray(nums1, nums2, i)
	// fmt.Println(conds)
	var a = 0
	if i-1 < 0 {
		a = nums2[j-1]
	} else if j-1 < 0 {
		a = nums1[i-1]
	} else {
		a = max2Ints(nums1[i-1], nums2[j-1])
	}
	if (m+n)%2 == 0 {
		var b = 0
		if i >= m {
			b = nums2[j]
		} else if j >= n {
			b = nums1[i]
		} else {
			b = min2Ints(nums1[i], nums2[j])
		}
		return float64(a+b) / 2
	} else {
		return float64(a)
	}
}
func convert(s string, numRows int) string {
	mem := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		mem[i] = make([]byte, 0)
	}
	const dBottom = true
	const dTop = false
	direct := dTop // 向下
	pos := 0
	for i := 0; i < len(s); i++ {
		mem[pos] = append(mem[pos], s[i])
		if pos == numRows-1 {
			direct = !direct
		} else if pos == 0 {
			direct = !direct
		}
		if numRows > 1 {
			if direct {
				pos++
			} else {
				pos--
			}
		}

	}
	ss := make([]string, numRows)
	for i, bs := range mem {
		ss[i] = string(bs)
	}
	return strings.Join(ss, "")
}
func reverse(x int) int {
	rev := 0
	maxMod := (math.MaxInt32 % 10)
	minMod := (math.MinInt32 % 10)
	// fmt.Println(math.MaxInt32 / 10)
	for x != 0 {
		pop := x % 10
		x /= 10
		// 214748364
		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop == maxMod) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop == minMod) {
			return 0
		}
		rev = rev*10 + pop
	}
	fmt.Println(rev)
	return rev
}
func myAtoi(s string) int {
	state := 0 // 0 space, 1 sign, 2 numbers, 3 other
	sign := 1
	n := 0
	maxMod := (math.MaxInt32 % 10)
	minMod := (math.MinInt32 % 10)
	for i := 0; i < len(s); i++ {
		if state == 0 {
			if s[i] == ' ' {
				continue
			} else if s[i] == '-' {
				sign = -1
				state = 2
			} else if s[i] == '+' {
				sign = 1
				state = 2
			} else if s[i] >= '0' && s[i] <= '9' {
				n = int(s[i]-'0') * sign
				state = 2
			} else {
				break
			}
		} else if state == 2 {
			if s[i] >= '0' && s[i] <= '9' {
				d := int(s[i] - '0')
				if n > math.MaxInt32/10 || (n == math.MaxInt32/10 && d >= maxMod) {
					n = math.MaxInt32
					return n
				}
				if n < math.MinInt32/10 || (n == math.MinInt32/10 && sign*d <= minMod) {
					n = math.MinInt32
					return n
				}
				n = n*10 + sign*d
			} else {
				break
			}
		}
	}
	return n
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x < 10 && x > 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	rev := 0
	for {
		// pop right
		r := x % 10
		x /= 10

		rev = rev*10 + r
		if rev == x {
			return true
		}
		if rev == x/10 {
			return true
		}
		if x < rev {
			return false
		}
	}
}
