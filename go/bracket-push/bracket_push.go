package brackets

const testVersion = 3

type stack []rune

func (s stack) push(x rune) stack {
	return append(s, x)
}

func (s stack) pop() stack {
	return s[:len(s)-1]
}

func (s stack) top() rune {
	return s[len(s)-1]
}

func (s stack) empty() bool {
	return len(s) == 0
}

var pair = map[rune]rune{'(': ')', '[': ']', '{': '}'}

func Bracket(s string) (bool, error) {
	st := make(stack, 0)

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			st = st.push(v)
		} else {
			if st.empty() || pair[st.top()] != v {
				return false, nil
			}
			st = st.pop()
		}
	}
	return st.empty(), nil
}
