package ch2

import "fmt"

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

//IncreamentLongPrecisionNumber incrments a number of any length
func IncreamentLongPrecisionNumber(in string) string {
	carry := rune(0)
	out := ""
	in = reverse(in)
	for i, v := range in {
		num := v - '0'
		if i == 0 {
			num = num + 1
		} else {
			num = num + carry
		}
		carry = num / 10
		num = num % 10
		out += fmt.Sprintf("%d", num)
	}
	if carry != 0 {
		out += fmt.Sprintf("%d", 1)
	}
	return reverse(out)
}
