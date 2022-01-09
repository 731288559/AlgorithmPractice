package medium

// 1218. 最长定差子序列

func longestSubsequence(arr []int, difference int) int {
	maxLen := 0
	m := make(map[int]int)
	for i := len(arr) - 1; i >= 0; i-- {
		m[arr[i]] = m[arr[i]+difference] + 1
		if m[arr[i]] > maxLen {
			maxLen = m[arr[i]]
		}
	}
	return maxLen
}

func longestSubsequence_v2(arr []int, difference int) (ans int) {
	dp := map[int]int{}
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if dp[v] > ans {
			ans = dp[v]
		}
	}
	return
}

func T_LC1218() {
	println(longestSubsequence([]int{1, 2, 3, 4}, 1), 4)
	println(longestSubsequence([]int{1, 3, 5, 7}, 1), 1)
	println(longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2), 4)
}
