package leetcode

import (
	"strings"
	"unicode"
)

var _Vowels = []byte{
	'A', 'E', 'I', 'O', 'U',
	'a', 'e', 'i', 'o', 'u',
}

func sortVowels(s string) string {
	m := make(map[byte]int, 10)
	var idxList []int
	for i, ch := range s {
		c := unicode.ToLower(ch)
		if strings.ContainsRune("aeiou", c) {
			m[s[i]]++
			idxList = append(idxList, i)
		}
	}

	ans := []byte(s)
	for _, idx := range idxList {
		var vowel byte
		for _, v := range _Vowels {
			cnt, ok := m[v]
			if ok && cnt > 0 {
				vowel = v
				m[v]--
				break
			}
		}
		ans[idx] = vowel
	}

	return string(ans)
}
