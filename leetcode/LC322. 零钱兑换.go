package leetcode

import (
	"math"
	"sort"
)

/*
	给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。

	计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。

	你可以认为每种硬币的数量是无限的。
*/

// 贪心版本
func coinChangeV1(coins []int, amount int) int {
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] >= coins[j]
	})

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		minTemp := math.MaxInt
		for _, coin := range coins {
			left := i - coin
			if left < 0 {
				continue
			}
			if dp[left] == math.MaxInt {
				continue
			}
			minTemp = min(minTemp, 1+dp[left])
		}
		dp[i] = minTemp
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

// DFS版本
func coinChangeV2(coins []int, amount int) int {
	n := len(coins)

	var dfs func(i, amount int) int

	dfs = func(i, amount int) int {
		if i < 0 {
			if amount == 0 {
				return 0
			}
			return math.MaxInt / 2
		}
		// 当前币值比目标值大，选不了
		if coins[i] > amount {
			return dfs(i-1, amount)
		}
		// 选择或者不选择当前硬币
		return min(dfs(i, amount-coins[i])+1, dfs(i-1, amount))
	}

	ans := dfs(n-1, amount)
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

// 递推版本
func coinChangeV3(coins []int, amount int) int {
	n := len(coins)

	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, amount+1)
	}
	for j := range f[0] {
		f[0][j] = math.MaxInt / 2
	}
	f[0][0] = 0
	for i, coin := range coins {
		for c := 0; c <= amount; c++ {
			if c < coin {
				f[i+1][c] = f[i][c]
			} else {
				f[i+1][c] = min(f[i][c], f[i+1][c-coin]+1)
			}
		}
	}

	ans := f[n][amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

// 滚动数组版本
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for j := range f {
		f[j] = math.MaxInt / 2
	}
	f[0] = 0
	for _, coin := range coins {
		for c := coin; c <= amount; c++ {
			f[c] = min(f[c], f[c-coin]+1)
		}
	}

	ans := f[amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}
