package medium

// 5201. 给植物浇水

func wateringPlants(plants []int, capacity int) int {
	curN := 0
	step := 0
	for i, n := range plants {
		if curN+n > capacity {
			step += 2 * i
			curN = n
		} else {
			curN += n
		}
		step++

	}
	return step
}

func T_LC5201() {
	println(wateringPlants([]int{2, 2, 3, 3}, 5), 14)
	println(wateringPlants([]int{1, 1, 1, 4, 2, 3}, 4), 30)
	println(wateringPlants([]int{7, 7, 7, 7, 7, 7, 7}, 8), 49)
}
