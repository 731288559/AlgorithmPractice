package medium

// 397. 整数替换

func integerReplacement(n int) int {
	cnt := 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			if n&2 == 2 && n != 3 {
				n++
			} else {
				n--
			}
		}
		cnt++
	}
	return cnt
}

func T_LC397() {
	// println(integerReplacement(8), 3)
	// println(integerReplacement(7), 4)
	// println(integerReplacement(2), 1)
	// println(integerReplacement(15), 5)
	println(integerReplacement(3), 2)
	// println(integerReplacement(10000), 16)
	println(integerReplacement(100000000), 31)
}
