package iteration

import "strings"

func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

func RepeatLoop(s string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
