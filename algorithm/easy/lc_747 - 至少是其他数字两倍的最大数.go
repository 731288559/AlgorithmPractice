package easy

func dominantIndex(nums []int) int {
	var index int
	var flag bool = true
	var maxNum int = nums[0]
	for i, n := range nums {
		if i == 0 {
			continue
		}

		if n > maxNum {
			if n >= 2*maxNum {
				flag = true
			} else {
				flag = false
			}
			maxNum = n
			index = i
		} else {
			if maxNum < n*2 {
				flag = false
			}
		}
	}

	if flag {
		return index
	}
	return -1
}

func T_LC747() {
	println(dominantIndex([]int{3, 6, 1, 0}), 1)
	println(dominantIndex([]int{1, 2, 3, 4}), -1)
	println(dominantIndex([]int{1}), 0)
	println(dominantIndex([]int{0, 0, 3, 2}), -1)
}
