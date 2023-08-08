package leetcode

/*
	1749. 任意子数组和的绝对值的最大值

	给你一个整数数组 nums 。一个子数组 [numsl, numsl+1, ..., numsr-1, numsr] 的 和的绝对值 为 abs(numsl + numsl+1 + ... + numsr-1 + numsr) 。

	请你找出 nums 中 和的绝对值 最大的任意子数组（可能为空），并返回该 最大值 。

	abs(x) 定义如下：

	如果 x 是负整数，那么 abs(x) = -x 。
	如果 x 是非负整数，那么 abs(x) = x 。
*/

func maxAbsoluteSum(nums []int) int {
	ans := 0

	abs := func(n int) int {
		if n >= 0 {
			return n
		}
		return -n
	}

	mayBenefit := func(a, b int) bool {
		if a == 0 || b == 0 {
			return true
		}
		if a*b > 0 {
			return true
		}
		return abs(b) <= abs(a)
	}

	tmp := 0
	for _, n := range nums {
		if mayBenefit(tmp, n) {
			tmp += n
		} else {
			tmp = n
		}
		if abs(tmp) > ans {
			ans = abs(tmp)
		}
	}

	return ans
}
