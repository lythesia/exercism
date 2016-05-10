package ocr

import (
	"strings"
)

const (
	zero = `
 _ 
| |
|_|
`
	one = `
   
  |
  |
`
	two = `
 _ 
 _|
|_ 
`
	three = `
 _ 
 _|
 _|
`
	four = `
   
|_|
  |
`
	five = `
 _ 
|_ 
 _|
`
	six = `
 _ 
|_ 
|_|
`
	seven = `
 _ 
  |
  |
`
	eight = `
 _ 
|_|
|_|
`
	nine = `
 _ 
|_|
 _|
`
)

var digit = map[string]string{
	zero: "0", one: "1", two: "2", three: "3", four: "4", five: "5", six: "6", seven: "7", eight: "8", nine: "9",
}

func recognizeDigit(s string) string {
	digit, ok := digit["\n"+s]
	if ok {
		return digit
	}
	return "?"
}

func Recognize(s string) []string {
	var numbers = []string{}
	var lineSeg []string
	var n int
	for i, e := range strings.Split(s, "\n") {
		if i%4 == 0 {
			// line ready
			lineNum := ""
			for _, seg := range lineSeg {
				lineNum += recognizeDigit(seg)
			}
			if lineNum != "" {
				numbers = append(numbers, lineNum)
			}
			continue
		}
		if i%4 == 1 {
			n = len(e) / 3
			lineSeg = make([]string, n)
		}
		for i := 0; i < n; i++ {
			lineSeg[i] += e[i*3:(i+1)*3] + "\n"
		}
	}
	return numbers
}
