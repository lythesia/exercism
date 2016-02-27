package house

import (
	"fmt"
	"strings"
)

func Embed(rel, noun string) string {
	return rel + " " + noun
}

func Verse(sub string, rel []string, noun string) string {
	if n := len(rel); n == 0 {
		return Embed(sub, noun)
	} else {
		return Verse(sub, rel[0:n-1], Embed(rel[n-1], noun))
	}
}

var arr = []struct{ noun, rel string }{
	{"house that Jack built.", ""},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

func verse(i int) string {
	ans := []string{"This is the " + arr[i].noun}
	for j := i; j > 0; j-- {
		ans = append(ans, fmt.Sprintf("that %s the %s", arr[j].rel, arr[j-1].noun))
	}
	return strings.Join(ans, "\n")
}

func Song() string {
	var verses []string
	for i := range arr {
		verses = append(verses, verse(i))
	}
	return strings.Join(verses, "\n\n")
}
