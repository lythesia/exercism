// +build !example

package foodchain

const testVersion = 2

// cow -> goat -> dog -> cat -> bird -> spider -> fly
var chain = [...]string{
	"",
	"She swallowed the spider to catch the fly.\n",
	"She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.\n",
	"She swallowed the cat to catch the bird.\n",
	"She swallowed the dog to catch the cat.\n",
	"She swallowed the goat to catch the dog.\n",
	"She swallowed the cow to catch the goat.\n",
}

var animal = [...]string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}

var how = [...]string{
	"",
	"It wriggled and jiggled and tickled inside her.\n",
	"How absurd to swallow a bird!\n",
	"Imagine that, to swallow a cat!\n",
	"What a hog, to swallow a dog!\n",
	"Just opened her throat and swallowed a goat!\n",
	"I don't know how she swallowed a cow!\n",
}

func refrain(i int) string {
	if i == 0 {
		return ""
	}
	return chain[i] + refrain(i-1)
}

// verse para x
func Verse(i int) string {
	ans := "I know an old lady who swallowed a " + animal[i-1] + ".\n"

	if i < 8 {
		ans += how[i-1] + refrain(i-1) + "I don't know why she swallowed the fly. Perhaps she'll die."
	} else {
		ans += "She's dead, of course!"
	}

	return ans
}

// verse para x to y
func Verses(i, j int) string {
	if i == j {
		return Verse(i)
	}
	return Verse(i) + "\n\n" + Verses(i+1, j)
}

// verse all
func Song() string {
	return Verses(1, 8)
}
