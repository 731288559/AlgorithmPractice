package leetcode

// 560. 和为 K 的子数组

// out of memory，本质上就是暴力，和dp关系不大
func subarraySum0(nums []int, k int) int {
	ans := 0

	n := len(nums)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = nums[i]
			} else {
				dp[i][j] = dp[i][j-1] + nums[j]
			}
			if dp[i][j] == k {
				ans++
			}
		}
	}

	return ans
}

func subarraySum(nums []int, k int) int {
	n := len(nums)
	l := make([]int, n+1)
	for i := range n {
		l[i+1] = l[i] + nums[i]
	}

	ans := 0
	cnt := make(map[int]int, n)
	for _, sum := range l {
		ans += cnt[sum-k] // sum - x = k, x = sum - k
		cnt[sum]++
	}
	return ans
}
