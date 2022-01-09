package easy

// 28. 实现 strStr()

func strStr(haystack string, needle string) int {
	// return strings.Index(haystack, needle)

	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func TestLC28() {
	println(strStr("hello", "ll") == 2)
	println(strStr("hehehe", "he") == 0)
	println(strStr("", "") == -1)
}
