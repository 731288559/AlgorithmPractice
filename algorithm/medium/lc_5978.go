package medium

func wordCount(startWords []string, targetWords []string) int {
	startMap := make(map[int]struct{})
	for _, word := range startWords {
		n := 0
		for _, char := range word {
			n |= 1 << (char - 'a')
		}
		startMap[n] = struct{}{}
	}

	var res int
	for _, word := range targetWords {
		sumN := 0
		for _, char := range word {
			sumN ^= 1 << (char - 'a')
		}
		for _, char := range word {
			if _, ok := startMap[sumN^1<<(char-'a')]; ok {
				res += 1
				break
			}
		}
	}

	return res
}

func T_LC5978() {
	println(wordCount([]string{"ant", "act", "tack"}, []string{"tack", "act", "acti"}), 2)
	println(wordCount([]string{"ab", "a"}, []string{"abc", "abcd"}), 1)
}
