package medium

import "strconv"

// 38. 外观数列

func countAndSay2(n int) string {
	if n == 1 {
		return "1"
	}

	str := countAndSay(n - 1)
	result := ""

	counter := 0
	var cur byte
	cur = str[0]
	for i := 0; i < len(str); i++ {
		if str[i] == cur {
			counter++
		} else {
			result = result + strconv.Itoa(counter) + string(cur)
			counter = 1
			cur = str[i]
		}
	}
	result = result + strconv.Itoa(counter) + string(cur)
	return result
}

func countAndSay(n int) string {
	result := "1"
	str := "1"
	for i := 1; i < n; i++ {
		var counter int
		var cur byte
		result = ""

		cur = str[0]
		for i := 0; i < len(str); i++ {
			if str[i] == cur {
				counter++
			} else {
				result = result + strconv.Itoa(counter) + string(cur)
				counter = 1
				cur = str[i]
			}
		}
		result = result + strconv.Itoa(counter) + string(cur)
		str = result
	}
	return result
}

func T_LC38() {
	println(countAndSay(1))
	println(countAndSay(2))
	println(countAndSay(3))
	println(countAndSay(4))
	println(countAndSay(5))
}
