package main

import (
	"reflect"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"/home/"}, "/home"},
		{"2", args{"/../"}, "/"},
		{"3", args{"/home//foo/"}, "/home/foo"},
		{"4", args{"/a/./b/../../c/"}, "/c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"2", args{[][]int{{1, 4}, {4, 5}}}, [][]int{{1, 5}}},
		{"t", args{[][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}}, [][]int{{2, 4}, {5, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_insert(t *testing.T) {
// 	type args struct {
// 		intervals   [][]int
// 		newInterval []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want [][]int
// 	}{
// 		{"1", args{[][]int{{1, 3}, {6, 9}}, []int{2, 5}}, [][]int{{1, 5}, {6, 9}}},
// 		{"0", args{[][]int{}, []int{5, 7}}, [][]int{{5, 7}}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := insert(tt.args.intervals, tt.args.newInterval); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("insert() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"12"}, 2},
		{"1", args{"226"}, 3},
		{"1", args{"0"}, 0},
		{"06", args{"06"}, 0},
		{"2", args{"8008"}, 0},
		{"t", args{"2101"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"1", args{"25525511135"}, []string{"255.255.11.135", "255.255.111.35"}},
		{"2", args{"0000"}, []string{"0.0.0.0"}},
		{"3", args{"1111"}, []string{"1.1.1.1"}},
		{"4", args{"010010"}, []string{"0.10.0.10", "0.100.1.0"}},
		{"5", args{"101023"}, []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
		{"t", args{"0279245587303"}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
