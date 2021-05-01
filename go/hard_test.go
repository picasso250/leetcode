package main

import (
	"reflect"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"1", args{1}, [][]string{{"Q"}}},
		{"2", args{2}, [][]string{}},
		{"4", args{4}, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"0", args{[]string{"a"}, 2}, []string{"a "}},
		{"0", args{[]string{"foo", "bar"}, 4}, []string{"foo ", "bar "}},
		{"0", args{[]string{"foo", "bar"}, 8}, []string{"foo bar "}},
		{"1", args{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16}, []string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		}},
		{"0", args{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16}, []string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		}},
		{"0", args{[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain",
			"to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20}, []string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evaluate(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"mult", args{"(mult 3 (add 2 3))"}, 15},
		{"outer", args{"(let x 1 y 2 x (add x y) (add x y))"}, 5},
		{"1", args{"(mult 3 (add 2 3))"}, 15},
		{"1", args{"(mult 3 (add 2 3))"}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluate(tt.args.expression); got != tt.want {
				t.Errorf("evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}
