package algs

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	for _, c := range []struct {
		input  string
		expect []string
	}{
		{"2", []string{"a", "b", "c"}},
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	} {
		got := LetterCombinations(c.input)
		if !reflect.DeepEqual(got, c.expect) {
			t.Errorf("input: %s\t|\texpect: %v\t|\tgot: %v", c.input, c.expect, got)
		}
	}
}
