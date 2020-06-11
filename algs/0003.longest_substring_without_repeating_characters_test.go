package algs

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	cs := []struct {
		input string
		want  int
	}{
		{
			input: "abcabcbb",
			want:  3,
		},
		{
			input: "bb",
			want:  1,
		},
		{
			input: "bbbbb",
			want:  1,
		},
		{
			input: "pwwkew",
			want:  3,
		},
		{
			input: " ",
			want:  1,
		},
		{
			input: "a",
			want:  1,
		},
		{
			input: "au",
			want:  2,
		},
		{
			input: "dvdf",
			want:  3,
		},
		{
			input: "ohomm",
			want:  3,
		},
		{
			input: "anviaj",
			want:  5,
		},
	}

	for _, c := range cs {
		got := LengthOfLongestSubstring(c.input)
		if got != c.want {
			t.Errorf("input: %q\twant: %d\tgot: %d\n", c.input, c.want, got)
		}
	}
}
