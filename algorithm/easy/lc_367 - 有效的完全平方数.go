package easy

// 367. 有效的完全平方数

func isPerfectSquare(num int) bool {
	l, r := 0, num
	var m int
	for l <= r {
		m = (l + r) / 2
		// println("l, r, m = ", l, r, m)
		if m*m == num {
			return true
		} else if m*m > num {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return false
}

func T_LC367() {
	println(isPerfectSquare(16))
	println(isPerfectSquare(14))
	println(isPerfectSquare(1))
}
