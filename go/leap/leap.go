package leap

const TestVersion = 1

func IsLeapYear(x int) bool {
	return (x%100 > 0 && x%4 == 0) || (x%400 == 0)
}
