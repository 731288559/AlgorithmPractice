package easy

import "math"

// 414. 第三大的数

func thirdMax(nums []int) int {
	a, b, c := math.MinInt64, math.MinInt64, math.MinInt64
	for _, i := range nums {
		if i > a {
			a, b, c = i, a, b
		} else if i < a {
			if i > b {
				b, c = i, b
			} else if i < b {
				if i > c {
					c = i
				}
			}
		}
	}

	if c == math.MinInt64 {
		return a
	}
	return c
}

func T_LC414() {
	nums := []int{3, 2, 1}
	println(thirdMax(nums))
	nums = []int{1, 2}
	println(thirdMax(nums))
	nums = []int{2, 2, 3, 1}
	println(thirdMax(nums))
	println(math.MinInt)
}
