package leetcode

func getSneakyNumbers(nums []int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[nums[i]] != nums[i] {
			t := nums[i]
			nums[i], nums[t] = nums[t], nums[i]
		}
	}
	return nums[n-2:]
}
