package medium

// 869. 重新排序得到 2 的幂

func reorderedPowerOf2(n int) bool {
	m := make(map[[10]int]struct{})

	for i := 1; i < 1e9; i = i << 1 {
		m[getCnts(i)] = struct{}{}
	}
	_, ok := m[getCnts(n)]
	return ok
}

func getCnts(n int) (cnt [10]int) {
	for n > 0 {
		cnt[n%10]++
		n = n / 10
	}
	return cnt
}

func ifPowerOf2(n int) bool {
	// if n <= 0 {
	// 	return false
	// }
	// counter := 0
	// for i := 0; i < 32; i++ {
	// 	if n>>i&1 == 1 {
	// 		counter++
	// 	}
	// 	if counter > 1 {
	// 		return false
	// 	}
	// }
	// if counter == 1 {
	// 	return true
	// }
	// return false
	return n > 0 && n&(n-1) == 0
}

func T_LC869() {
	println(reorderedPowerOf2(1), true)
	println(reorderedPowerOf2(10), false)
	println(reorderedPowerOf2(16), true)
	println(reorderedPowerOf2(24), false)
	println(reorderedPowerOf2(46), true)
}
