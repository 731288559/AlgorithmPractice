package easy

// 575. 分糖果

func distributeCandies(candyType []int) int {
	m := make(map[int]struct{})
	for _, i := range candyType {
		m[i] = struct{}{}
	}
	n := len(candyType)
	if len(m) > n/2 {
		return n / 2
	}
	return len(m)
}

func T_LC575() {
	println(distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	println(distributeCandies([]int{1, 1, 2, 3}))
	// println(distributeCandies([]int{3, 3, 2, 1, 1, 2}))
}
