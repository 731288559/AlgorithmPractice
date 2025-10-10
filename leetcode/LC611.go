package leetcode

import (
	"slices"
	"sort"
)

// 611. 有效三角形的个数

// 暴力解法
func triangleNumber_1(nums []int) int {
	sort.Ints(nums)

	ans := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j] <= nums[k] {
					break
				}
				ans++
			}
		}
	}
	return ans
}

func triangleNumber(nums []int) int {
	slices.Sort(nums)

	ans := 0
	n := len(nums)
	for k := 2; k < n; k++ {
		i := 0
		j := k - 1
		c := nums[k]
		for i < j {
			a, b := nums[i], nums[j]
			if a+b > c {
				ans += j - i
				j--
			} else {
				i++
			}
		}
	}
	return ans
}
