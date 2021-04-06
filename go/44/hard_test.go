package main

import "testing"

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
		{"1", args{"", ""}, true},
		{"1", args{"", "a"}, false},
		{"1", args{"", "*"}, true},
		{"1", args{"b", "a"}, false},
		{"1", args{"baa", "b*"}, true},
		{"1", args{"baa", "*a"}, true},
		{"1", args{"aa", "a"}, false},
		{"2", args{"aa", "*"}, true},
		{"3", args{"cb", "?a"}, false},
		{"4", args{"adceb", "*a*b"}, true},
		{"5", args{"acdcb", "a*c?b"}, false},
		{"t1", args{"aab", "c*a*b"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
