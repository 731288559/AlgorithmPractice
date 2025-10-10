package leetcode

import "math"

func maximumEnergy(energy []int, k int) int {
	ans := math.MinInt

	n := len(energy)
	for i := 0; i < k; i++ {
		idx := n - 1 - i
		tempMax := math.MinInt
		curSum := 0
		for idx >= 0 {
			curSum += energy[idx]
			tempMax = max(curSum, tempMax)
			idx -= k
		}
		ans = max(ans, tempMax)
	}

	return ans
}
