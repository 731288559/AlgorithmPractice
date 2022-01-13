package hard

// 335. 路径交叉

func isSelfCrossing(distance []int) bool {
	n := len(distance)
	if n < 4 {
		return false
	}
	for i := 3; i < n; i++ {
		if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
			return true
		}
		if i >= 4 {
			if distance[i-1] == distance[i-3] && distance[i] >= distance[i-2]-distance[i-4] {
				return true
			}
		}
		if i >= 5 {
			if distance[i] >= distance[i-2]-distance[i-4] &&
				distance[i-1] >= distance[i-3]-distance[i-5] &&
				distance[i-2] > distance[i-4] &&
				distance[i-1] <= distance[i-3] {
				return true
			}
		}
		// // 第 1 类路径交叉的情况
		// if distance[i] >= distance[i-2] && distance[i-1] <= distance[i-3] {
		//     return true
		// }

		// // 第 2 类路径交叉的情况
		// if i == 4 && distance[3] == distance[1] &&
		//     distance[4] >= distance[2]-distance[0] {
		//     return true
		// }

		// // 第 3 类路径交叉的情况
		// if i >= 5 && distance[i-3]-distance[i-5] <= distance[i-1] &&
		//     distance[i-1] <= distance[i-3] &&
		//     distance[i] >= distance[i-2]-distance[i-4] &&
		//     distance[i-2] > distance[i-4] {
		//     return true
		// }
	}

	return false
}

func T_LC335() {
	println(isSelfCrossing([]int{2, 1, 1, 2}), true)
	println(isSelfCrossing([]int{1, 2, 3, 4}), false)
	println(isSelfCrossing([]int{1, 1, 1, 1}), true)
}
