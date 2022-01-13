package medium

func minSwaps(nums []int) int {
	oneNum := 0
	for i := range nums {
		if nums[i] == 1 {
			oneNum += 1
		}
	}

	var oneNumMax int
	oneNumTmp := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			for j := 0; j < oneNum; j++ {
				if nums[j%len(nums)] == 1 {
					oneNumTmp += 1
				}
			}
		} else {
			oneNumTmp = oneNumTmp - nums[i-1] + nums[(i+oneNum-1)%len(nums)]
		}

		if oneNumTmp > oneNumMax {
			oneNumMax = oneNumTmp
		}
	}

	return oneNum - oneNumMax
}

func T_LC5977() {
	println(minSwaps([]int{0, 1, 0, 1, 1, 0, 0}), 1)
	println(minSwaps([]int{0, 1, 1, 1, 0, 0, 1, 1, 0}), 2)
	println(minSwaps([]int{1, 1, 0, 0, 1}), 0)
}
