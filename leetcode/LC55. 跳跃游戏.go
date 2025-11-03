package leetcode

func canJump(nums []int) bool {
	mx := 0
	for i := range nums {
		if i > mx {
			return false
		}

		mx = max(mx, i+nums[i])
		if mx > len(nums)-1 {
			return true
		}
	}

	return true
}
