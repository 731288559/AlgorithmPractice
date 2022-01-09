package medium

func lastRemaining(n int) int {
	res := 1
	step := 1
	fromLeft := true

	for n > 1 {
		if fromLeft || n%2 == 1 {
			res += step
		}
		step <<= 1
		n >>= 1
		fromLeft = !fromLeft
	}
	return res
}

func T_LC390() {
	println(lastRemaining(9))
	println(lastRemaining(1))
}
