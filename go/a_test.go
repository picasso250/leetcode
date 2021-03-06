package main

import (
	"reflect"
	"testing"
)

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"abcde", "ace"}, 3},
		{"2", args{"abc", "abc"}, 3},
		{"3", args{"abcde", "def"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"abcabcbb"}, 3},
		{"2", args{"bbbbb"}, 1},
		{"3", args{"pwwkew"}, 3},
		{"4", args{"aa"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"0", args{[]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8, 10}}, 5.5},
		{"1", args{[]int{1, 3}, []int{2}}, 2.0},
		{"2", args{[]int{1, 2}, []int{3, 4}}, 2.5},
		{"3", args{[]int{0, 0}, []int{0, 0}}, 0.0},
		{"4", args{[]int{}, []int{1}}, 1.0},
		{"5", args{[]int{2}, []int{}}, 2.0},
		{"6", args{[]int{0, 0, 0, 0, 0}, []int{-1, 0, 0, 0, 0, 0, 1}}, 0.0}, //[0,0,0,0,0] [-1,0,0,0,0,0,1]
		{"7", args{[]int{}, []int{2, 3}}, 2.5},                              //[] [2,3]
		{"8", args{[]int{3}, []int{-2, -1}}, -1.0},                          //[3] [-2,-1]
		{"9", args{[]int{4, 5}, []int{1, 2, 3}}, 3.0},                       //[4,5] [1,2,3]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"2", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{"3", args{"A", 1}, "A"},
		{"4", args{"AB", 1}, "AB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{123}, 321},
		{"2", args{-123}, -321},
		{"3", args{120}, 21},
		{"4", args{0}, 0},
		{"5", args{214748364}, 463847412},
		{"6", args{463847412}, 214748364},
		{"6", args{4638474127}, 0},
		{"6", args{4638474128}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"42"}, 42},
		{"1", args{"   -42"}, -42},
		{"1", args{"4193 with words"}, 4193},
		{"1", args{"words and 987"}, 0},
		{"1", args{"-91283472332"}, -2147483648},
		{"1", args{"-2147483649"}, -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1122", args{1122}, false},
		{"121", args{121}, true},
		{"-121", args{-121}, false},
		{"10", args{10}, false},
		{"-101", args{-101}, false},
		{"1", args{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// {"empty", args{"", ""}, true},
		// {"a", args{"a", "a"}, true},
		// {".", args{"a", "."}, true},
		{".*", args{"a", ".*"}, true},
		{"1", args{"aa", "a"}, false},
		{"2", args{"aa", "a*"}, true},
		{"3", args{"ab", ".*"}, true},
		{"4", args{"aab", "c*a*b"}, true},
		{"5", args{"mississippi", "mis*is*p*."}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1x1", args{[]int{1, 1}}, 1},
		{"1", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"3", args{[]int{4, 3, 2, 1, 4}}, 16},
		{"4", args{[]int{1, 2, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"I", args{"I"}, 1},
		{"III", args{"III"}, 3},
		{"IV", args{"IV"}, 4},
		{"IX", args{"IX"}, 9},
		{"LVIII", args{"LVIII"}, 58},
		{"MCMXCIV", args{"MCMXCIV"}, 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty", args{[]string{}}, ""},
		{"empty string", args{[]string{"a", ""}}, ""},
		{"1", args{[]string{"flower", "flow", "flight"}}, "fl"},
		{"2", args{[]string{"dog", "racecar", "car"}}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"2", args{[]int{}}, [][]int{}},
		{"3", args{[]int{0}}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"23"}, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"empty", args{""}, []string{}},
		{"3", args{"2"}, []string{"a", "b", "c"}},
		{"t", args{"234"}, []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinations(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{toList([]int{1, 2, 3, 4, 5}), 2}, toList([]int{1, 2, 3, 5})},
		{"2", args{toList([]int{1}), 1}, toList([]int{})},
		{"3", args{toList([]int{1, 2}), 1}, toList([]int{1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); got.ToString() != tt.want.ToString() {
				t.Errorf("removeNthFromEnd() = %s, want %s", got.ToString(), tt.want.ToString())
			}
		})
	}
}

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"()"}, true},
		{"2", args{"()[]{}"}, true},
		{"3", args{"(]"}, false},
		{"4", args{"([)]"}, false},
		{"5", args{"{[]}"}, true},
		{"6", args{"{[()]}"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{toList([]int{1, 2, 4}), toList([]int{1, 3, 4})}, toList([]int{1, 1, 2, 3, 4, 4})},
		{"2", args{toList([]int{}), toList([]int{})}, toList([]int{})},
		{"3", args{toList([]int{}), toList([]int{0})}, toList([]int{0})},
		{"t1", args{toList([]int{-9, 3}), toList([]int{5, 7})}, toList([]int{-9, 3, 5, 7})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got.ToString(), tt.want.ToString()) {
				t.Errorf("mergeTwoLists() = %s, want %s", got.ToString(), tt.want.ToString())
			}
		})
	}
}

func Test_generateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"3", args{3}, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"1", args{1}, []string{"()"}},
		{"2", args{2}, []string{"(())", "()()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{[]*ListNode{
			toList([]int{1, 4, 5}),
			toList([]int{1, 3, 4}),
			toList([]int{2, 6}),
		}}, toList([]int{1, 1, 2, 3, 4, 4, 5, 6})},
		{"2", args{[]*ListNode{}}, toList([]int{})},
		{"3", args{[]*ListNode{toList([]int{})}}, toList([]int{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{toList([]int{1, 2, 3, 4})}, toList([]int{2, 1, 4, 3})},
		{"2", args{toList([]int{})}, toList([]int{})},
		{"3", args{toList([]int{1})}, toList([]int{1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got.ToString(), tt.want.ToString()) {
				t.Errorf("swapPairs() = %s, want %s", got.ToString(), tt.want.ToString())
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	type result struct {
		length int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want result
	}{
		{"1", args{[]int{1, 1, 2}}, result{2, []int{1, 2}}},
		{"2", args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, result{5, []int{0, 1, 2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); !reflect.DeepEqual(result{got, tt.args.nums[:got]}, tt.want) {
				t.Errorf("removeDuplicates() = %v %v, want %v", got, tt.args.nums[:got], tt.want)
			}
		})
	}
}

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{10, 3}, 3},
		{"2", args{7, -3}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{"barfoothefoobarman", []string{"foo", "bar"}}, []int{0, 9}},
		{"2", args{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}}, []int{}},
		{"t1", args{"barfoofoobarthefoobarman", []string{"bar", "foo", "the"}}, []int{6, 9, 12}},
		{"t2", args{
			"lingmindraboofooowingdingbarrwingmonkeypoundcake",
			[]string{"fooo", "barr", "wing", "ding", "wing"}}, []int{13}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{"()"}, 2},
		{"1", args{"(()"}, 2},
		{"2", args{")()())"}, 4},
		{"3", args{""}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}}, true},
		{"2", args{[][]byte{
			{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudoku(tt.args.board); got != tt.want {
				t.Errorf("isValidSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{[][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(tt.args.board)
			solveSudokuPrint(tt.args.board)
			if !isValidSudoku(tt.args.board) {
				t.Errorf("not isValidSudoku() ")
			}
		})
	}
}

func Test_solveSudokuCollectHoles(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantRet int
	}{
		{"1", args{[][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}}, 51},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := solveSudokuCollectHoles(tt.args.board); !reflect.DeepEqual(len(gotRet), tt.wantRet) {
				t.Errorf("solveSudokuCollectHoles() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"2", "3"}, "6"},
		{"2", args{"123", "456"}, "56088"},
		{"0", args{"0", "0"}, "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clumsy(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 1},
		{"2", args{2}, 2 * 1},
		{"3", args{3}, 3 * 2 / 1},
		{"1", args{4}, 4*3/2 + 1},
		{"2", args{10}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clumsy(tt.args.N); got != tt.want {
				t.Errorf("clumsy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"1", args{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}}, [][]string{
			{"ate", "eat", "tea"},
			{"nat", "tan"},
			{"bat"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"1", args{2.0, 10}, 1024.0},
		{"2", args{2.1, 3}, 9.26100},
		{"3", args{2.0, -2}, 0.25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
