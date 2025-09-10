package leetcode

func canPartition(nums []int) bool {
	sum := 0
	n := len(nums)
	maxNum := 0
	for _, num := range nums {
		sum += num
		maxNum = max(maxNum, num)
	}

	if sum%2 > 0 {
		return false
	}

	target := sum / 2
	if maxNum > target {
		return false
	}

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	dp[1][nums[0]] = true

	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			k := nums[i-1]
			if j < k {
				// 无法选择
				dp[i][j] = dp[i-1][j]
				continue
			}
			// 可选可不选
			dp[i][j] = dp[i-1][j] || dp[i-1][j-k]

		}
	}

	return dp[n][target]
}
