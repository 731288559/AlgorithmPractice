package leetcode

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
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
