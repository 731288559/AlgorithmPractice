package medium

// 3. 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	var (
		maxLen int
		left   int
		m      map[byte]int
	)
	m = make(map[byte]int)
	for idx := range s {
		if _, ok := m[s[idx]]; ok {
			if m[s[idx]]+1 > left {
				left = m[s[idx]] + 1
			}
		}
		m[s[idx]] = idx
		if idx-left+1 > maxLen {
			maxLen = idx - left + 1
		}
	}
	return maxLen
}

func TestLC3() {
	println(lengthOfLongestSubstring("abcabcbb") == 3) // 3
	println(lengthOfLongestSubstring("bbbbbbb") == 1)  // 1
	println(lengthOfLongestSubstring("pwwkew") == 3)   // 3
	println(lengthOfLongestSubstring("abba") == 2)     // 2
}
