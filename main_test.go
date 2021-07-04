package main

import (
	"testing"
)

func TestFun(t *testing.T) {
	in := ")()())"
	out := longestValidParentheses(in)
	want := 4
	if out != want {
		t.Errorf("got %d, want %d", out, want)
	}
}
