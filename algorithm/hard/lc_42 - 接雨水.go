package hard

// 42. 接雨水

func trap(height []int) int {
	// n := len(height)
	// left := make([]int, n)
	// right := make([]int, n)

	// max := 0
	// for i := 0; i < n; i++ {
	// 	if height[i] > max {
	// 		max = height[i]
	// 	}
	// 	left[i] = max
	// }
	// max = 0
	// for i := n - 1; i >= 0; i-- {
	// 	if height[i] > max {
	// 		max = height[i]
	// 	}
	// 	right[i] = max
	// }

	// var result int
	// for i := range height {
	// 	result += minInt(left[i], right[i]) - height[i]
	// }
	// return result
	n := len(height)

	var result int
	var leftMax, rightMax int
	left, right := 0, n-1
	for left < right {
		leftMax = maxInt(leftMax, height[left])
		rightMax = maxInt(rightMax, height[right])
		if leftMax < rightMax {
			result += leftMax - height[left]
			left++
		} else {
			result += rightMax - height[right]
			right--
		}
	}

	return result
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func T_LC42() {
	println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	println(trap([]int{4, 2, 0, 3, 2, 5}))
}
