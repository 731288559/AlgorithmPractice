package easy

import (
	"fmt"
	"strings"
)

// 500. 键盘行

func findWords(words []string) []string {
	m := make(map[rune]int)
	for _, i := range "qwertyuiop" {
		m[i] = 0
	}
	for _, i := range "asdfghjkl" {
		m[i] = 1
	}
	for _, i := range "zxcvbnm" {
		m[i] = 2
	}

	var result []string
	for _, word := range words {
		f := -1
		ok := true
		for _, i := range strings.ToLower(word) {
			if f < 0 {
				f = m[i]
			} else {
				if m[i] != f {
					ok = false
					break
				}
			}
		}
		if ok {
			result = append(result, word)
		}
	}
	return result
}

func T_LC500() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
	fmt.Println(findWords([]string{"omk"}))
	fmt.Println(findWords([]string{"adsdf", "sfd"}))
}
