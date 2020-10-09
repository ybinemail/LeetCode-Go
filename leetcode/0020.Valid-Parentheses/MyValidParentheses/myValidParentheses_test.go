package MyValidParentheses

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case1",args{"()"},true},
		{"case2",args{"()[]{}"},true},
		{"case3",args{"(]"},false},
		{"case4",args{"([)]"},false},
		{"case5",args{"{[]}"},true},
		{"case6",args{"{"},false},
		{"case7",args{"{{"},false},
		{"case8",args{"{}}"},false},
		{"case9",args{"{{}}"},true},
		{"case10",args{""},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
