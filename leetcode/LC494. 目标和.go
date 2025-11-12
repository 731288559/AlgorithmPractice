package leetcode

/*
	给你一个非负整数数组 nums 和一个整数 target 。

	向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

	例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
	返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
*/

func zeroOneKnapsack(capacity int, weights []int, values []int) int {
	n := len(weights)

	var dfs func(i int, c int) int
	dfs = func(i int, c int) int {
		if i < 0 {
			return 0
		}
		if c < weights[i] {
			return dfs(i-1, c)
		}
		return max(dfs(i-1, c), dfs(i-1, c-weights[i]+values[i]))
	}

	return dfs(n-1, capacity)
}

func findTargetSumWays(nums []int, target int) int {
	// 设正数的和为 p
	// 则负数的和为 sum - p （绝对值）
	// p - (sum - p) = target
	// => 2p = target + sum
	// => p = (target + sum) / 2
	// 问题等价于从nums选一些数，使得和恰好等于 (sum+target)/2 的方案数
	for _, n := range nums {
		target += n
	}
	// 由于p必定是非负整数，因此2p必定是非负偶数，即target+sum必定是非负偶数
	if target < 0 || target%2 == 1 {
		return 0
	}
	target /= 2

	n := len(nums)
	var dfs func(i int, c int) int
	dfs = func(i int, c int) int {
		if i < 0 {
			if c == 0 {
				return 1
			}
			return 0
		}
		if c < nums[i] {
			return dfs(i-1, c)
		}
		return dfs(i-1, c) + dfs(i-1, c-nums[i])
	}

	return dfs(n-1, target)
}
