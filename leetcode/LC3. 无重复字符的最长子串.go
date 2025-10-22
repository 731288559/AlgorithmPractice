package leetcode

func lengthOfLongestSubstring(s string) int {
	ans := 0

	m := make(map[rune]int)
	tempLen := 0
	start := -1
	for i, c := range s {
		idx, ok := m[c]
		if !ok || idx < start {
			tempLen++
		} else {
			tempLen = i - idx
			start = idx
		}
		m[c] = i
		ans = max(ans, tempLen)
	}

	return ans
}
