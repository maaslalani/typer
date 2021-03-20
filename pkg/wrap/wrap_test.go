package wrap

import (
	"testing"
)

func TestWrap(t *testing.T) {
	cases := []struct {
		Input, Output string
		Width         int
	}{
		{"foo", "foo", 4},
		{"foo bar", "foo\nbar", 4},
		{"foofoo barbar", "foofoo\nbarbar", 4},
		{"foobar foobar foobar", "foobar\nfoobar\nfoobar", 4},
		{"foo bar baz", "foo\nbar\nbaz", 1},
		{"foo bar baz", "foo bar baz", 10},
		{"foo bar, foo bar, foo bar", "foo bar,\nfoo bar,\nfoo bar", 6},
		{"foo bar, foo bar, foo bar", "foo bar, foo bar, foo bar", 60},
		{"foo bar,\nfoo bar,\nfoo bar", "foo bar,\nfoo bar,\nfoo bar", 60},
	}

	for _, tc := range cases {
		got := Words(tc.Input, tc.Width)
		if got != tc.Output {
			t.Fatalf("Input:\n%s\n\nWanted:\n%s\n\nGot:\n%s", tc.Input, tc.Output, got)
		}
	}
}
