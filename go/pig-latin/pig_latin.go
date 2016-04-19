package igpay

import "strings"

var consonant = []string{
	"ch",
	"qu",
	"ph",
	"sch",
	"sh",
	"squ",
	"thr",
	"th",
	"wh",
	"zh",
	"wr",
}

func isVowel(r rune) bool {
	return strings.ContainsRune("aeiou", r)
}

func beginWithVowel(s string) bool {
	if isVowel(rune(s[0])) {
		return true
	} else if (rune(s[0]) == 'x' || rune(s[0]) == 'y') && !isVowel(rune(s[1])) {
		return true
	}
	return false
}

func PigLatinWord(word string) string {
	if beginWithVowel(word) {
		return word + "ay"
	}
	for _, p := range consonant {
		if strings.HasPrefix(word, p) {
			return word[len(p):] + p + "ay"
		}
	}
	return word[1:] + string(word[0]) + "ay"
}

func PigLatin(s string) string {
	words := strings.Split(s, " ")
	ans := make([]string, len(words))
	for i, e := range words {
		ans[i] = PigLatinWord(e)
	}
	return strings.Join(ans, " ")
}
