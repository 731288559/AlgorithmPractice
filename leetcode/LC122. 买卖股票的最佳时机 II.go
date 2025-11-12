package leetcode

import "math"

/*
	给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

	在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。然而，你可以在 同一天 多次买卖该股票，但要确保你持有的股票不超过一股。

	返回 你能获得的 最大 利润 。
*/

// 无记忆化，会超时
func maxProfitV1(prices []int) int {
	n := len(prices)

	var dfs func(i int, hold bool) int
	dfs = func(i int, hold bool) int {
		if i < 0 {
			if hold {
				return math.MinInt
			}
			return 0
		}
		if hold {
			// 持有情况下什么也不做 / 前一天买入
			return max(dfs(i-1, true), dfs(i-1, false)-prices[i])
		}
		// 未持有情况下什么也不做 / 前一天卖出
		return max(dfs(i-1, false), dfs(i-1, true)+prices[i])
	}

	return dfs(n-1, false)
}

func maxProfitV2(prices []int) int {
	n := len(prices)

	dp := make([][2]int, n+1)
	dp[0][1] = math.MinInt

	for i := range prices {
		dp[i+1][0] = max(dp[i][0], dp[i][1]+prices[i])
		dp[i+1][1] = max(dp[i][1], dp[i][0]-prices[i])
	}

	return dp[n][0]
}

// 滚动数组，O（1）空间复杂度
func maxProfit(prices []int) int {
	f0 := 0
	f1 := math.MinInt
	newf0 := 0
	for i := range prices {
		newf0 = max(f0, f1+prices[i])
		f1 = max(f1, f0-prices[i])
		f0 = newf0
	}

	return f0
}
