package phonenumber

import (
	"fmt"
	"unicode"
)

func extract(raw string) string {
	ans := make([]rune, 0, len(raw))
	for _, e := range raw {
		if unicode.IsDigit(e) {
			ans = append(ans, e)
		}
	}
	return string(ans)
}

func Number(numberStr string) (string, error) {
	numberStr = extract(numberStr)
	n := len(numberStr)
	if n == 10 {
		return numberStr, nil
	} else if n == 11 && numberStr[0] == '1' {
		return numberStr[1:], nil
	} else {
		return "", fmt.Errorf("invalid phone number.")
	}
}

func AreaCode(numberStr string) (string, error) {
	if number, err := Number(numberStr); err != nil {
		return "", err
	} else {
		return number[0:3], nil
	}
}

func Format(numberStr string) (string, error) {
	if number, err := Number(numberStr); err != nil {
		return "", err
	} else {
		area, _ := AreaCode(number)
		return fmt.Sprintf("(%s) %s-%s", area, number[3:6], number[len(number)-4:]), nil
	}
}
