package leetcode

// 300. 最长递增子序列

func lengthOfLIS(nums []int) int {
	ans := 0

	n := len(nums)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
		ans = max(ans, dp[j])
	}

	return ans
}
