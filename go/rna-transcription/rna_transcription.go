package strand

const testVersion = 3

var tab = map[rune]rune{'C': 'G', 'G': 'C', 'T': 'A', 'A': 'U'}

func ToRNA(s string) string {
	ans := make([]rune, len(s))
	for i, e := range s {
		ans[i] = tab[e]
	}
	return string(ans)
}
