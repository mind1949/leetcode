package algs

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	for _, c := range []struct {
		haystack string
		needle   string
		expect   int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
		{"aabaaabaaac", "aabaaac", 4},
	} {
		got := StrStr(c.haystack, c.needle)
		if got != c.expect {
			format := "haystack: %s\t|\tneedle: %s\t|\texpect: %d\t|\tgot: %d"
			t.Errorf(format, c.haystack, c.needle, c.expect, got)
		}
	}
}

func TestStrStrRabinKarp(t *testing.T) {
	for _, c := range []struct {
		haystack string
		needle   string
		expect   int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
		{"aabaaabaaac", "aabaaac", 4},
	} {
		got := StrStrRabinKarp(c.haystack, c.needle)
		if got != c.expect {
			format := "haystack: %s\t|\tneedle: %s\t|\texpect: %d\t|\tgot: %d"
			t.Errorf(format, c.haystack, c.needle, c.expect, got)
		}
	}
}
