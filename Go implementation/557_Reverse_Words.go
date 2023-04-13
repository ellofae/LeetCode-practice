package main

import (
	"strings"
)

func reverseWords(s string) string {
	arr := make([]byte, 0)

	for _, str := range strings.Split(s, " ") {
		for j := len(str) - 1; j >= 0; j-- {
			arr = append(arr, str[j])
		}
		arr = append(arr, ' ')
	}
	return string(arr)
}

func main() {
	reverseWords("Let's take LeetCode contest")

}
