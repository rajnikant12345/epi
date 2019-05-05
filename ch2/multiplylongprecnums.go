package ch2

import (
	"fmt"
	"strings"
)

//MultiplyLongs multiply two long integers
func MultiplyLongs(one string, two string) string {
	one = reverse(one)
	two = reverse(two)
	result := make([]int, len(one)+len(two))

	for i := 0; i < len(two); i++ {
		for j := 0; j < len(one); j++ {
			res := (one[j] - '0') * (two[i] - '0')
			result[i+j] = (result[i+j] + int(res))
		}
	}
	carry := 0
	num := 0
	var out string
	for i := 0; i < len(result); i++ {
		num = carry + result[i]
		carry = num / 10
		num = num % 10
		out += fmt.Sprintf("%d", num)
	}
	if carry != 0 {
		out += fmt.Sprintf("%d", num)
	}
	return strings.TrimLeft(reverse(out), "0")
}
