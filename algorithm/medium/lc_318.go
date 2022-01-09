package medium

// 318. 最大单词长度乘积

func maxProduct_0(words []string) int {
	masks := make([]int, len(words))
	for i, n1 := range words {
		for _, n2 := range n1 {
			masks[i] |= 1 << (n2 - 'a')
		}
	}
	ans := 0
	for i, n1 := range words {
		for j, n2 := range words {
			if masks[i]&masks[j] == 0 && len(n1)*len(n2) > ans {
				ans = len(n1) * len(n2)
			}
		}
	}
	return ans
}

// 优化：相同字母构成的单词，仅记录长度最长的一项
func maxProduct(words []string) int {
	ans := 0
	masks := make(map[int]int)
	for _, n := range words {
		mask := 0
		for _, n2 := range n {
			mask |= 1 << (n2 - 'a')
		}
		if len(n) > masks[mask] {
			masks[mask] = len(n)
		}
	}
	for i, n := range masks {
		for j, n2 := range masks {
			if i&j == 0 && n*n2 > ans {
				ans = n * n2
			}
		}
	}
	return ans
}

func T_LC318() {
	println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}), 16)
	println(maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}), 4)
	println(maxProduct([]string{"a", "aa", "aaa", "aaaa"}), 0)
}
