package main

import (
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
