package easy

import (
	"sort"
)

// 1005. K 次取反后最大化的数组和

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	sum := 0
	min := 100
	for _, n := range nums {
		if k == 0 {
			sum += n
			continue
		}
		if n >= 0 {
			if n < min {
				min = n
			}
		} else {
			n = -n
			min = n
			k--
		}
		sum += n
	}
	if k > 0 && k%2 > 0 {
		sum -= 2 * min
	}
	return sum
}

func T_LC1005() {
	println(largestSumAfterKNegations([]int{4, 2, 3}, 1), 5)
	println(largestSumAfterKNegations([]int{3, -1, 0, 2}, 3), 6)
	println(largestSumAfterKNegations([]int{2, -3, -1, 5, -4}, 2), 13)
	println(largestSumAfterKNegations([]int{-8, 3, -5, -3, -5, -2}, 6), 22)
	println(largestSumAfterKNegations([]int{-2, -2, 0, 2, 5}, 3), 11)
	println(largestSumAfterKNegations([]int{8, -7, -3, -9, 1, 9, -6, -9, 3}, 8), 53)

}
