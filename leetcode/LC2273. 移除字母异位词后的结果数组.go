package leetcode

import (
	"bytes"
	"slices"
)

func removeAnagrams(words []string) []string {
	k := 0

	var last []byte
	for _, word := range words {
		s := []byte(word)
		slices.Sort(s)

		if !bytes.Equal(last, s) {
			words[k] = word
			k++
			last = s
		}
	}

	return words[:k]
}
