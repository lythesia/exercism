package counter

import (
	"testing"
)

var tests = []struct {
	input                 string
	lines, letters, chars int
}{
	{"one line\n", 1, 7, 9},
	{"two\nlines", 2, 8, 9},
	{"ends in \\n\n", 1, 7, 11},
	{"", 0, 0, 0},
	{"this is 平仮名\n", 1, 9, 12},
}

func TestLines(t *testing.T) {
	for _, test := range tests {
		c := makeCounter()
		c.AddString(test.input)
		if actual := c.Lines(); actual != test.lines {
			t.Fatalf("Lines(%q) = %d, expected %d.", test.input, actual, test.lines)
		}
	}
}

func TestLetters(t *testing.T) {
	for _, test := range tests {
		c := makeCounter()
		c.AddString(test.input)
		if actual := c.Letters(); actual != test.letters {
			t.Fatalf("Letters(%q) = %d, expected %d.", test.input, actual, test.letters)
		}
	}
}

func TestChracters(t *testing.T) {
	for _, test := range tests {
		c := makeCounter()
		c.AddString(test.input)
		if actual := c.Characters(); actual != test.chars {
			t.Fatalf("Characters(%q) = %d, expected %d.", test.input, actual, test.chars)
		}
	}
}
