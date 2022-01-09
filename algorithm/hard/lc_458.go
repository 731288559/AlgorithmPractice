package hard

import "math"

// 458. 可怜的小猪
// 一只猪能够获得 轮次数+1 的信息

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	k := minutesToTest / minutesToDie
	// 1 ping -> k+1 status
	// (k+1)^n >= buckets
	// 换底公式 log以a为底b为指数 = log以c为底b为指数 / log以c为底a为指数
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(k+1))))
}

func T_LC458() {
	println(poorPigs(1000, 15, 60), 5)
	println(poorPigs(4, 15, 15), 2)
	println(poorPigs(4, 15, 30), 2)
}
