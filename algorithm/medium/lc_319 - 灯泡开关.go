package medium

// 319. 灯泡开关

func bulbSwitch(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		if i*i > n {
			break
		}
		ans++
	}
	return ans
	// return int(math.Sqrt(float64(n) + 0.5))
}

func T_LC319() {
	println(bulbSwitch(3), 1)
	println(bulbSwitch(0), 0)
	println(bulbSwitch(1), 1)
	println(bulbSwitch(9), 3)
}
