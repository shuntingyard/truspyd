/*
 * The example to get going with unit tests.
 */
package truspyd

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		has, want string
	}{
		{"Hello, new gopher", "rehpog wen ,olleH"},
		{"Hallöchen...", "...nehcöllaH"},
	}
	for _, c := range tests {
		got := Reverse(c.has)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.has, got, c.want)
		}
	}
}
