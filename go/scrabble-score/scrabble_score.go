package scrabble

import "unicode"

const testVersion = 3

func toScore(ch rune) int {
	switch unicode.ToUpper(ch) {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	}
	return 0
}

func Score(s string) int {
	ans := 0
	for _, e := range s {
		ans += toScore(e)
	}
	return ans
}
