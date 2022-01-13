package easy

// 520. 检测大写字母

func detectCapitalUse(word string) bool {
	if len(word) <= 1 {
		return true
	}
	if word[0] >= 'a' && word[0] <= 'z' {
		for i := range word {
			if word[i] < 'a' || word[i] > 'z' {
				return false
			}
		}
		return true
	}
	if word[1] >= 'A' && word[1] <= 'Z' {
		for i := 2; i < len(word); i++ {
			if word[i] < 'A' || word[i] > 'Z' {
				return false
			}
		}
		return true
	} else {
		for i := 2; i < len(word); i++ {
			if word[i] < 'a' || word[i] > 'z' {
				return false
			}
		}
		return true
	}
}

func T_LC520() {
	println(detectCapitalUse("leetcode"), true)
	println(detectCapitalUse("USA"), true)
	println(detectCapitalUse("FlaG"), false)
}
