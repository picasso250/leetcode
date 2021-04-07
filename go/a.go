package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
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
func matches(s, p byte) bool {
	if p == '.' {
		return true
	}
	return s == p
}
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
		dp[i][0] = true
	}
	dp[0][0] = true
	for i := 0; i < len(dp); i++ {
		for j := 0; j <= len(p); j++ {
			// if i == 0 && j == 0 {
			// 	continue
			// }
			if i-1 >= 0 && j-1 >= 0 && p[j-1] != '*' {
				if matches(s[i-1], p[j-1]) {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = false
				}
			} else if i-1 >= 0 && j-2 >= 0 {
				if matches(s[i-1], p[j-2]) {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
func maxArea(height []int) int {
	max := -1
	i := 0
	j := len(height) - 1
	for i < j {
		a := j - i
		h := min2Ints(height[i], height[j])
		s := a * h
		if s > max {
			max = s
		}
		if height[i] < height[j] { // 移动小的
			i++
		} else {
			j--
		}
	}
	return max
}

type roman struct {
	name  string
	value int
}

func build9(rm10 roman, rm1 roman) roman {
	return roman{rm1.name + rm10.name, rm10.value - rm1.value}
}
func build4(rm5 roman, rm1 roman) roman {
	return roman{rm1.name + rm5.name, rm5.value - rm1.value}
}
func buildN(rm roman, n int) roman {
	s, v := "", 0
	for i := 0; i < n; i++ {
		s += rm.name
		v += rm.value
	}
	return roman{s, v}
}

func romanToInt(s string) int {
	// const (
	// 	i = iota //b=0
	// 	v        //c=1
	// 	x        //c=1
	// 	l        //c=1
	// 	c        //c=1
	// 	d        //c=1
	// 	m        //c=1
	// )
	// om := []roman{
	// 	{"I", 1},
	// 	{"V", 5},
	// 	{"X", 10},
	// 	{"L", 50},
	// 	{"C", 100},
	// 	{"D", 500},
	// 	{"M", 1000},
	// }
	// table := []roman{om[m]}
	// table = append(table, build9(om[m], om[c]))
	// table = append(table, om[d])
	// table = append(table, build4(om[d], om[c]))
	// table = append(table, buildN(om[c], 3))
	// table = append(table, buildN(om[c], 2))
	// table = append(table, om[c])
	// table = append(table, build9(om[c], om[x]))
	// table = append(table, om[l])
	// table = append(table, build4(om[l], om[x]))
	// table = append(table, buildN(om[x], 3))
	// table = append(table, buildN(om[x], 2))
	// table = append(table, om[x])
	// table = append(table, build9(om[x], om[i]))
	// table = append(table, om[v])
	// table = append(table, build4(om[v], om[i]))
	// table = append(table, buildN(om[i], 3))
	// table = append(table, buildN(om[i], 2))
	// table = append(table, om[i])
	// for i := len(table) - 1; i >= 0; i-- {
	// 	fmt.Printf("%v\n", table[i])
	// }
	// fmt.Printf("%v\n", table)
	table := []roman{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"CCC", 300},
		{"CC", 200},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"XXX", 30},
		{"XX", 20},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"III", 3},
		{"II", 2},
		{"I", 1},
	}
	tryMatch := func(start int) (length int, value int) {
		for _, r := range table {
			length = indexStrStart(s, start, r.name)
			if length > 0 {
				return length, r.value
			}
		}
		return 0, 0
	}
	i := 0
	value := 0
	for i < len(s) {
		l, v := tryMatch(i)
		value += v
		i += l
	}
	return value
}
func indexStrStart(s string, start int, pat string) (length int) {
	if start+len(pat) > len(s) {
		return -1
	}
	for j := 0; j < len(pat); j++ {
		if s[start+j] != pat[j] {
			return -1
		}
	}
	return len(pat)
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	length := minLenStrs(strs)
	if length == 0 {
		return ""
	}
	common := make([]byte, length)
	for i := 0; i < length; i++ {
		for j, s := range strs {
			if len(s) == 0 {
				return ""
			}
			if j == 0 {
				common[i] = s[i]
			} else {
				if common[i] != s[i] {
					return string(common[:i])
				}
			}
		}
	}
	return string(common)
}
func minLenStrs(strs []string) (length int) {
	length = len(strs[0])
	for _, s := range strs {
		if len(s) < length {
			length = len(s)
		}
	}
	return length
}
func threeSum(nums []int) [][]int {
	m := make(map[int]int)
	ns := make([]int, len(nums))
	copy(ns, nums)
	sort.Ints(ns)
	for i, n := range ns {
		m[n] = i
	}
	r := make(map[string][]int)
	for i := 0; i < len(ns); i++ {
		for j := len(ns) - 1; j > i; j-- {
			c := 0 - ns[i] - ns[j]
			ii, ok := m[c]
			if ok && c >= ns[j] && (ii != i && ii != j) {
				A := []int{ns[i], ns[j], c}
				// sort.Ints(A)
				S := []string{strconv.Itoa(A[0]), strconv.Itoa(A[1]), strconv.Itoa(A[2])}
				r[strings.Join(S, ",")] = A
			}
		}
	}
	ret := make([][]int, 0)
	for _, x := range r {
		ret = append(ret, x)
	}
	return ret
}
func letterCombinations(digits string) []string {
	m := map[byte][]byte{
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}
	ret := make([][]byte, 0)
	for i := 0; i < len(digits); i++ {
		bs := m[digits[i]]
		// map
		if len(ret) == 0 {
			for _, b := range bs {
				ret = append(ret, []byte{b})
			}
		} else {
			r := make([][]byte, 0)
			for _, rr := range ret {
				for _, b := range bs {
					rrr := make([]byte, len(rr)+1)
					copy(rrr[:len(rr)], rr)
					rrr[len(rr)] = b
					r = append(r, rrr)
				}
			}
			ret = r
		}
	}
	rr := make([]string, len(ret))
	for i, r := range ret {
		rr[i] = string(r)
	}
	return rr
}

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	var prevDel *ListNode = nil
	var toDel *ListNode = nil
	var prev *ListNode
	i := 0
	for p != nil {
		if i >= n-1 {
			prevDel = prev
		}
		if i >= n {
			toDel = p
		}
		i++
		prev = p
		p = p.Next
	}
	// del
	head = head.delNode(toDel, prevDel)
	return head
}
func (head *ListNode) delNode(toDel *ListNode, prevDel *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if toDel == head {
		newHead := head.Next
		head.Next = nil
		return newHead
	}
	if prevDel != nil {
		if prevDel.Next != toDel {
			panic("prevDel's next is not toDel")
		}
		prevDel.Next = toDel.Next
		toDel.Next = nil
	}
	return head
}
func toList(lst []int) *ListNode {
	if len(lst) == 0 {
		return nil
	}
	l := make([]ListNode, len(lst))
	for i, v := range lst {
		var next *ListNode
		if i == len(lst)-1 {
			next = nil
		} else {
			next = &l[i+1]
		}
		l[i] = ListNode{v, next}
	}
	return &l[0]
}
func (head *ListNode) toSlice() []int {
	a := make([]int, 0)
	for p := head; p != nil; p = p.Next {
		a = append(a, p.Val)
	}
	return a
}
func (head *ListNode) toStrings() []string {
	a := make([]string, 0)
	for p := head; p != nil; p = p.Next {
		a = append(a, strconv.Itoa(p.Val))
	}
	return a
}
func (head *ListNode) ToString() string {
	return fmt.Sprintf("[%s]", strings.Join(head.toStrings(), ","))
}
func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if c == ')' {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if c == ']' {
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else if c == '}' {
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val < l2.Val {
		head = l1
	} else {
		head = l2
	}
	var p *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			// push l1
			t := l1.Next
			l1.Next = nil
			p = pushListNode(p, l1)
			l1 = t
		} else {
			t := l2.Next
			l2.Next = nil
			p = pushListNode(p, l2)
			l2 = t
		}
	}
	if l1 != nil {
		pushListNode(p, l1)
	}
	if l2 != nil {
		pushListNode(p, l2)
	}
	return head
}
func pushListNode(tail *ListNode, node *ListNode) *ListNode {
	if node == nil {
		panic("node is nil")
	}
	// if node.Next != nil {
	// 	panic("node.Next is not nil")
	// }
	if tail == nil {
		tail = node
	} else {
		tail.Next = node
	}
	return node
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}
	// 划分
	ret := []string{}
	for i := 0; i <= n-1; i++ {
		j := n - 1 - i
		left, right := generateParenthesis(i), generateParenthesis(j)
		for _, ll := range left {
			for _, rr := range right {
				ret = append(ret, "("+ll+")"+rr)
			}
		}
	}
	sort.Strings(ret)
	return ret
}
func mergeKLists(lists []*ListNode) *ListNode {
	// k := len(lists)
	// ss := make([][]int, k)
	s := make([]int, 0)
	for _, l := range lists {
		// ss[i] = l.toSlice()
		s = append(s, toInts(l)...)
	}
	sort.Ints(s)
	return toList(s)
}
func toInts(head *ListNode) []int {
	a := make([]int, 0)
	for p := head; p != nil; p = p.Next {
		a = append(a, p.Val)
	}
	return a
}
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	prev := &ListNode{0, head}
	dummy := prev
	n1 := head
	for {
		prev, n1 = swap2(prev, n1, n1.Next)
		if n1 == nil {
			break
		}
	}
	return dummy.Next
}

// next.Next == tail
func swap2(prev *ListNode, n1 *ListNode, n2 *ListNode) (next *ListNode, tail *ListNode) {
	if n1 == nil {
		return nil, nil
	}
	if n2 == nil {
		return nil, nil
	}
	prev.Next = n2
	t := n2.Next
	n2.Next = n1
	n1.Next = t
	return n1, t
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	return head
}
func reverseList(dummy *ListNode) *ListNode {
	head := dummy.Next
	if head == nil {
		return nil
	}
	p := head
	var prev *ListNode
	var newHead *ListNode
	for {
		t := p.Next
		if prev != nil {
			p.Next = prev
		}
		if t == nil {
			newHead = p
			break
		}
		prev = p
		p = t
	}
	return newHead
}
func removeDuplicates(nums []int) int {
	d := 0
	prev := 0
	for i := 0; i < len(nums); i++ {
		if i > 0 {
			if nums[i] == prev {
				d++
			} else {
				nums[i-d] = nums[i]
			}
		}
		prev = nums[i]
	}
	return len(nums) - d
}
func removeElement(nums []int, val int) int {
	d := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			d++
		} else {
			nums[i-d] = nums[i]
		}
	}
	return len(nums) - d
}
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	left := dividend
	i := 0
	for left >= divisor {
		left -= divisor
		i++
	}
	if sign > 0 {
		return i
	} else {
		return -i
	}
}
func findSubstring(s string, words []string) []int {
	if len(s) == 0 {
		return []int{}
	}
	if len(words) == 0 {
		return []int{}
	}
	// n := wordsLength(words)
	wordMap := words2map(words)
	n := len(words)
	k := len(words[0])
	ret := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if findSubstringAllMatch(s[i:], n, k, wordMap) {
			ret = append(ret, i)
		} else {
		}
	}
	return ret
}
func words2map(words []string) map[string]int {
	m := map[string]int{}
	for _, word := range words {
		m[word]++
	}
	return m
}
func wordsLength(words []string) int {
	n := 0
	for _, word := range words {
		n += len(word)
	}
	return n
}
func findSubstringAllMatch(s string, n, k int, wordMap map[string]int) bool {
	if len(s) < n*k {
		return false
	}
	m := str2map(s[:n*k], k)
	return mapEqual(wordMap, m)
}
func str2map(s string, k int) map[string]int {
	m := map[string]int{}
	for i := 0; i < len(s); i += k {
		m[s[i:i+k]]++
	}
	return m
}
func mapEqual(a, b map[string]int) bool {
	for k := range a {
		if b[k] != a[k] {
			return false
		}
	}
	for k := range b {
		if a[k] != b[k] {
			return false
		}
	}
	return true
}
func removeWord(words []string, j int) []string {
	a := words[:j]
	b := words[j+1:]
	ret := make([]string, len(a)+len(b))
	copy(ret[:j], a)
	copy(ret[j:], b)
	return ret
}
func findSubStringBegin(s string, word string) bool {
	if len(s) < len(word) {
		return false
	}
	for i := 0; i < len(s) && i < len(word); i++ {
		if s[i] != word[i] {
			return false
		}
	}
	return true
}

type parPos struct {
	begin int
	end   int
}

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	posStack := make([]parPos, 0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '(' {
			stack = append(stack, i)
		} else if c == ')' {
			if len(stack) == 0 {
				// do nothing
			} else {
				pos := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				posStack = pushPosStack(posStack, pos, i+1)
			}
		}
	}
	return maxParsLen(posStack)
}
func maxParsLen(ps []parPos) int {
	max := 0
	for _, p := range ps {
		length := p.end - p.begin
		if length > max {
			max = length
		}
	}
	return max
}
func pushPosStack(s []parPos, begin, end int) []parPos {
	p := parPos{begin, end}
	if len(s) == 0 {
		return append(s, p)
	}
	last := s[len(s)-1]
	if inPar(last, p) {
		s[len(s)-1] = p
	} else {
		s = append(s, p)
	}
	if len(s) >= 2 {
		last = s[len(s)-1]
		last2 := s[len(s)-2]
		if isParConnect(last2, last) {
			s = append(s[:len(s)-2], parConnect(last2, last))
		}
	}
	return s
}
func inPar(small, big parPos) bool {
	return small.begin > big.begin && small.end < big.end
}
func isParConnect(a, b parPos) bool {
	return a.end == b.begin
}
func parConnect(a, b parPos) parPos {
	return parPos{a.begin, b.end}
}
func isValidSudoku(board [][]byte) bool {
	// 数字 1-9 在每一行只能出现一次。
	for i := 0; i < 9; i++ {
		m := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			n := board[i][j]
			if n != '.' {
				if m[n] {
					return false
				}
				m[n] = true
			}
		}
	}
	// 数字 1-9 在每一列只能出现一次。
	for i := 0; i < 9; i++ {
		m := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			n := board[j][i]
			if n != '.' {
				if m[n] {
					return false
				}
				m[n] = true
			}
		}
	}
	// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !isValid33(board, i*3, j*3) {
				return false
			}
		}
	}
	return true
}
func isValid33(board [][]byte, i int, j int) bool {
	m := make(map[byte]bool)
	for ii := 0; ii < 3; ii++ {
		for jj := 0; jj < 3; jj++ {
			n := board[i+ii][j+jj]
			if n != '.' {
				if m[n] {
					return false
				}
				m[n] = true
			}
		}
	}
	return true
}
func solveSudoku(board [][]byte) {
	var mi [9][10]bool
	var mj [9][10]bool
	var m3 [9][10]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			n := board[i][j]
			if n != '.' {
				nn := n - '0'
				mi[i][nn] = true
				mj[j][nn] = true
				m3[i/3*3+j/3][nn] = true
			}
		}
	}
	holes := solveSudokuCollectHoles(board)
	solveSudokuIter(board, mi, mj, m3, holes)
}
func solveSudokuCollectHoles(board [][]byte) (ret [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			n := board[i][j]
			if n == '.' {
				ret = append(ret, []int{i, j})
			}
		}
	}
	return
}
func solveSudokuIter(board [][]byte, mi, mj, m3 [9][10]bool, holes [][]int) bool {
	if len(holes) == 0 {
		return true
	}
	i, j := holes[0][0], holes[0][1]
	// fmt.Printf("for %d,%d\n", i, j)
	m := solveSudokuOccupy(mi, mj, m3, i, j)
	possible := solveSudokuPossible(m)
	if len(possible) == 0 {
		// solveSudokuPrint(board)
		// fmt.Printf("no possible solution for %d,%d\n", i, j)
		return false
	}
	for _, p := range possible {
		// fmt.Printf("try %d for %d,%d\n", p, i, j)
		board[i][j] = byte('0' + p)
		mi[i][p] = true
		mj[j][p] = true
		m3[i/3*3+j/3][p] = true
		if solveSudokuIter(board, mi, mj, m3, holes[1:]) {
			return true
		}
		mi[i][p] = false
		mj[j][p] = false
		m3[i/3*3+j/3][p] = false
		board[i][j] = '.'
	}
	// fmt.Printf("try %v for %d,%d, find nothing\n", possible, i, j)
	return false
}
func solveSudokuOccupy(mi, mj, m3 [9][10]bool, i, j int) (m [10]bool) { // true for not possible
	for k, v := range mi[i] {
		if v {
			m[k] = true
		}
	}
	for k, v := range mj[j] {
		if v {
			m[k] = true
		}
	}
	for k, v := range m3[i/3*3+j/3] {
		if v {
			m[k] = true
		}
	}
	return
}
func solveSudokuPrint(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Printf("\n")
	}
}
func solveSudokuPossible(m [10]bool) (ret []int) {
	for i := 1; i <= 9; i++ {
		if m[i] == false {
			ret = append(ret, i)
		}
	}
	return
}
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := []byte(num1), []byte(num2)
	var aa []byte
	zeros := []byte{}
	for i := len(n2) - 1; i >= 0; i-- {
		b := n2[i]
		a := multiply1(n1, b)
		a = append(a, zeros...)
		zeros = append(zeros, '0')
		if len(aa) == 0 {
			aa = a
		} else {
			aa = multiplyAdd(aa, a)
		}
	}
	return string(multiplyRemoveLeadingZeros(aa))
}
func multiplyRemoveLeadingZeros(a []byte) []byte {
	for i := 0; i < len(a); i++ {
		if a[i] != '0' {
			return a[i:]
		}
	}
	return []byte{'0'}
}
func multiply1(num1 []byte, num2 byte) []byte {
	size := len(num1)
	ret := make([]byte, size+1)
	n2 := num2 - '0'
	carry := 0
	for i := size - 1; i >= 0; i-- {
		n1 := num1[i] - '0'
		c := int(n1*n2) + carry
		if c >= 10 {
			carry = c / 10
			c = c % 10
		} else {
			carry = 0
		}
		ret[i+1] = byte('0' + c)
	}
	if carry > 0 {
		ret[0] = byte(carry + '0')
	} else {
		ret = ret[1:]
	}
	return ret
}
func multiplyAdd(num1, num2 []byte) []byte {
	carry := 0
	size1, size2 := len(num1), len(num2)
	if size1 > size2 {
		num2 = append(make([]byte, size1-size2), num2...)
	} else {
		num1 = append(make([]byte, size2-size1), num1...)
	}
	size := max2Ints(size1, size2)
	ret := make([]byte, size+1)
	for i := size - 1; i >= 0; i-- {
		a, b := multiplyByte2int(num1[i]), multiplyByte2int(num2[i])
		var n int
		n, carry = multiplyAdd3(a, b, carry)
		ret[i+1] = byte('0' + n)
	}
	if carry > 0 {
		ret[0] = byte('0' + carry)
	} else {
		ret = ret[1:]
	}
	return ret
}
func multiplyByte2int(a byte) int {
	if a == 0 {
		return 0
	}
	return int(a - '0')
}
func multiplyAdd3(a, b, c int) (n, carry int) {
	n = a + b + c
	if n > 9 {
		n = n - 10
		carry = 1
	}
	return
}
func clumsy(N int) int {
	n := N / 4
	if n*4 < N {
		n++
	}
	s := 0
	for i := 0; i < n; i++ {
		a, b, c, d := i*4+0, i*4+1, i*4+2, i*4+3
		a, b, c, d = N-a, N-b, N-c, N-d
		var sign int
		if i == 0 {
			sign = 1
		} else {
			sign = -1
		}
		var ss int
		if a == 1 {
			ss = sign * a
		} else if b == 1 {
			ss = sign * a * b
		} else if c == 1 {
			ss = sign * a * b / c
		} else {
			ss = sign*a*b/c + d
		}
		s += ss
	}
	return s
}
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {

	}
	if n%2 == 0 {

	} else {

	}
}
func rotateTopLeftToCenter(n, i, j int) (x, y int) {
	c := n / 2
	if n%2 == 0 {
		if i < 2 {
			x = i - c
		} else {
			x = i - c + 1
		}
		if j < 2 {
			x = -j + c
		} else {
			x = -j + c - 1
		}
		return
	} else {
		return i - c, -j + c
	}
}
func rotateCenterToTopLeft(n, x, y int) (i, j int) {
	return x + n/2, y + n/2
}
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		sb := []byte(s)
		asb := make([]byte, len(sb))
		copy(asb, sb)
		ByteSlice(asb).Sort()
		m[string(asb)] = append(m[string(asb)], s)
	}
	ret := make([][]string, 0)
	for _, ss := range m {
		ret = append(ret, ss)
	}
	return ret
}

// ByteSlice attaches the methods of Interface to []int, sorting in increasing order.
type ByteSlice []byte

func (x ByteSlice) Len() int           { return len(x) }
func (x ByteSlice) Less(i, j int) bool { return x[i] < x[j] }
func (x ByteSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x ByteSlice) Sort() { sort.Sort(x) }
