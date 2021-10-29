package atoi

import (
	"errors"
)

func Atoi(s string) (int, error) {
	runes := []rune(s)

	var n int
	isNegative := false
	for pos, char := range runes {
		if pos == 0 && char == '-' {
			isNegative = true
		} else if char >= '0' && char <= '9' {
			n = n * 10 + int(char - '0')
		} else {
			return 0, errors.New("string contains non-integer characters")
		}
	}

	if isNegative {
		n = -n
	}
	return int(n), nil
}
