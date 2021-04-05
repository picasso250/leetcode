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
