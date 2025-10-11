package leetcode

/*
	1793. 好子数组的最大分数

	给你一个整数数组 nums （下标从 0 开始）和一个整数 k 。
	一个子数组 (i, j) 的 分数 定义为 min(nums[i], nums[i+1], ..., nums[j]) * (j - i + 1) 。一个 好 子数组的两个端点下标需要满足 i <= k <= j 。
	请你返回 好 子数组的最大可能 分数 。

	示例 1：
	输入：nums = [1,4,3,7,4,5], k = 3
	输出：15
	解释：最优子数组的左右端点下标是 (1, 5) ，分数为 min(4,3,7,4,5) * (5-1+1) = 3 * 5 = 15 。

	示例 2：
	输入：nums = [5,5,4,5,4,1,1,1], k = 0
	输出：20
	解释：最优子数组的左右端点下标是 (0, 4) ，分数为 min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20 。
*/

/*
	本题要计算的分数，和 84. 柱状图中最大的矩形 是一样的，计算的是最大矩形面积。只不过多了一个约束：矩形必须包含下标 k
	从下标k开始做贪心，直到双指针遍历完整个数组
*/

func maximumScore(nums []int, k int) int {
	i, j := k, k
	ans := nums[k]
	minH := nums[k]
	n := len(nums)
	for t := 0; t < n-1; t++ {
		if j == n-1 || i > 0 && nums[i-1] > nums[j+1] {
			i--
			minH = min(minH, nums[i])
		} else {
			j++
			minH = min(minH, nums[j])
		}
		ans = max(ans, minH*(j-i+1))
	}

	return ans
}
