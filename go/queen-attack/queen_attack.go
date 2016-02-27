package queenattack

import (
	"errors"
	"fmt"
	"unicode"
)

func CanQueenAttack(w, b string) (bool, error) {
	if w == b {
		return false, errors.New("same square")
	}
	wp, e := getpos(w)
	if e != nil {
		return false, e
	}
	bp, e := getpos(b)
	if e != nil {
		return false, e
	}
	return wp.x == bp.x || wp.y == bp.y || abs(wp.x-bp.x) == abs(wp.y-bp.y), nil
}

type square struct{ x, y int }

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func getpos(s string) (square, error) {
	var p square
	if len(s) != 2 || !unicode.IsLetter(rune(s[0])) || !unicode.IsDigit(rune(s[1])) {
		return p, fmt.Errorf("not position format: %s", s)
	}

	p.x, p.y = int(s[0]-'a'), int(s[1]-'0')
	if p.x < 0 || p.x > 7 || p.y < 0 || p.y > 7 {
		return p, fmt.Errorf("out of board: %s", s)
	}

	return p, nil
}
