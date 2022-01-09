package medium

// 45. 跳跃游戏 II

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	var (
		curIdx int
		step   int
		maxIdx int
	)
	step++
	for curIdx+nums[curIdx] < len(nums)-1 {
		for i := curIdx; i <= curIdx+nums[curIdx]; i++ {
			if nums[i]+i > maxIdx+nums[maxIdx] {
				maxIdx = i
			}
		}
		if maxIdx == curIdx {
			return -1
		}
		curIdx = maxIdx
		step++
	}
	return step
}

func TestLC45() {
	println(jump([]int{2, 3, 1, 1, 4})) // 2
	println(jump([]int{2, 3, 0, 1, 4})) // 2
	println(jump([]int{1, 1, 1, 1, 1})) // 5

	println(jump([]int{1, 2, 0, 0, 0}))
}
