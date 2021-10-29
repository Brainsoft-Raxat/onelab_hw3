package itoa

import "strings"

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Itoa(n int) (s string) {
	isNegative := false

	var sb strings.Builder

	if n < 0 {
		isNegative = true
		n = -n
	}

	for {
		r := n % 10
		sb.WriteByte(byte(r)  + '0')
		n /= 10
		if n == 0 {
			break
		}
	}

	if isNegative {
		sb.WriteByte('-')
	}

	return reverse(sb.String())
}