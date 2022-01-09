package easy

// 598. 范围求和 II

func maxCount(m int, n int, ops [][]int) int {
	minA := 0
	minB := 0
	for _, op := range ops {
		if op[0] < minA || minA == 0 {
			minA = op[0]
		}
		if op[1] < minB || minB == 0 {
			minB = op[1]
		}
	}
	if minA == 0 || minB == 0 {
		return m * n
	}
	return minA * minB
}

func T_LC598() {
	println(maxCount(3, 3, [][]int{{2, 2}, {3, 3}}))
	println(maxCount(3, 3, [][]int{}))
}
