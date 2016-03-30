package allergies

var items = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(score int) []string {
	var ans = []string{}
	for i := 7; i >= 0; i-- {
		if t := 1 << uint(i); score >= t {
			n := score / t
			score -= n * t
			ans = append([]string{items[i]}, ans...)
		}
	}
	return ans
}

func AllergicTo(score int, item string) bool {
	for _, e := range Allergies(score) {
		if e == item {
			return true
		}
	}
	return false
}
