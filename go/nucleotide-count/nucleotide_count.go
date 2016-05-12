package dna

import (
	"fmt"
	"strings"
)

type DNA string
type Histogram map[rune]int

func (d DNA) Counts() Histogram {
	var count = Histogram{'A': 0, 'T': 0, 'C': 0, 'G': 0}
	for _, e := range d {
		count[e]++
	}
	return count
}

func (d DNA) Count(r byte) (int, error) {
	if strings.IndexRune(string(d), rune(r)) == -1 {
		return 0, fmt.Errorf("invalid nucleotide: %v", r)
	}
	count := d.Counts()
	return count[rune(r)], nil
}
