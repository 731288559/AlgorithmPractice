package leetcode

/*
	53. 最大子数组和

	给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
	子数组 是数组中的一个连续部分。
*/

func maxSubArray(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i], nums[i]+nums[i-1])
		ans = max(ans, nums[i])
	}
	return ans
}
