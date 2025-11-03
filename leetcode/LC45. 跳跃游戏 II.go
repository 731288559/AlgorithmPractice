package leetcode

func jump(nums []int) int {
	ans := 0

	n := len(nums)
	curRight := 0
	maxRight := 0
	for i := range nums {
		// 如果当前已经能到终点（最后一个位置），直接返回次数
		if curRight >= n-1 {
			return ans
		}
		// 无法走到当前位置，因此无法走到终点
		if i > curRight {
			return -1
		}

		// 目前能到达的最远位置
		maxRight = max(maxRight, i+nums[i])

		// 在当前不增加跳跃次数时候能走到的最远位置，如果到了，则跳往目前能达到的最远位置，次数加一
		if i == curRight {
			ans++
			curRight = maxRight
		}
	}

	return ans
}
