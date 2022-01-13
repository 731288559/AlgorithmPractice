package medium

import "math"

func increasingTriplet(nums []int) bool {
	first := nums[0]
	second := math.MaxInt
	for _, n := range nums {
		if n > second {
			return true
		}
		if n > first {
			second = n
		} else {
			first = n
		}
	}
	return false
}

func T_LC334() {
	println(increasingTriplet([]int{1, 2, 3, 4, 5}), true)
	println(increasingTriplet([]int{5, 4, 3, 2, 1}), false)
	println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}), true)
}
