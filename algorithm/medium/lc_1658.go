package medium

// 1658. 将 x 减到 0 的最小操作数

func minOperations(nums []int, x int) int {
	// return helper1658(nums, x, 0)
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum < x {
		return -1
	}
	if sum == x {
		return len(nums)
	}
	sum = sum - x

	cur := 0
	maxLen := -1
	left, right := 0, 0
	for right < len(nums) {
		cur += nums[right]
		right++
		// if cur < sum {
		// 	continue
		// }

		for left <= right && cur > sum {
			cur -= nums[left]
			left++
			// println("left++:", left)
		}

		if cur == sum {
			println("left,right:", left, right-1)
			if right-left > maxLen {
				maxLen = right - left
			}
		}
	}
	if maxLen < 0 {
		return -1
	}
	return len(nums) - maxLen
}

func T_LC1658() {
	println(minOperations([]int{1, 1, 4, 2, 3}, 5))
	// println(minOperations([]int{5, 6, 7, 8, 9}, 4))
	// println(minOperations([]int{3, 2, 20, 1, 1, 3}, 10))
	println(minOperations([]int{8828, 9581, 49, 9818, 9974, 9869, 9991, 10000, 10000, 10000, 9999, 9993, 9904, 8819, 1231, 6309}, 134365))
}
