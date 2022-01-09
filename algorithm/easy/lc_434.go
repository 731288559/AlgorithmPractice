package easy

// 434. 字符串中的单词数

func countSegments(s string) int {
	var result int
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if s[i] != ' ' {
				result++
			}
			continue
		}
		if s[i] != ' ' && s[i-1] == ' ' {
			result++
		}
	}
	return result
}

func T_LC434() {
	println(countSegments("Hello, my name is John"), 5)
	println(countSegments("aaa"), 1)
	println(countSegments("aaa  "), 1)
	println(countSegments("aaa   b"), 2)
	println(countSegments("Of all the gin joints in all the towns in all the world,   "), 13)
	println(countSegments("    "), 0)
}
