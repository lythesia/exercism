package beer

import "fmt"

const tmpl = "%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n"

func Verse(n int) (string, error) {
	if n > 99 {
		return "", fmt.Errorf("invalid verses: %d", n)
	}
	switch n {
	case 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	case 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	default:
		return fmt.Sprintf(tmpl, n, n, n-1), nil
	}
}

func Verses(upper, lower int) (string, error) {
	if upper > 99 || lower < 0 || upper < lower {
		return "", fmt.Errorf("invalid range: [%d, %d]", lower, upper)
	}
	verses := ""
	for i := upper; i >= lower; i-- {
		versei, _ := Verse(i)
		verses += versei + "\n"
	}
	return verses, nil
}

func Song() string {
	if song, ok := Verses(99, 0); ok == nil {
		return song
	}
	return ""
}
