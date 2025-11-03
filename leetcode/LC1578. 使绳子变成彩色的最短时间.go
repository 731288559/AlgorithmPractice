package leetcode

func minCost(colors string, neededTime []int) int {
	ans := 0

	n := len(colors)
	preColor := byte(0)
	preColorMax := -1
	for i := 0; i < n; i++ {
		// 和前一个同色
		if colors[i] == preColor {
			ans += min(preColorMax, neededTime[i])
			preColorMax = max(preColorMax, neededTime[i])
			continue
		}
		// 和前一个不同色
		preColor = colors[i]
		preColorMax = neededTime[i]
	}

	return ans
}
