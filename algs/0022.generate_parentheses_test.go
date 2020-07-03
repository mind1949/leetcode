package algs

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	for _, c := range []struct {
		input  int
		expect []string
	}{
		{0, nil},
		{1, []string{"()"}},
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	} {
		got := GenerateParenthesis(c.input)
		sort.Strings(got)
		sort.Strings(c.expect)
		if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("input: %d\t|\texpect: %+v\t|\tgot: %+v", c.input, c.expect, got)
		}
	}
}
