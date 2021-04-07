package main

import "testing"

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
