package easy

// 852. 山脉数组的峰顶索引

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l < r {
		m := (l + r) / 2
		if arr[m] > arr[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return r
}

func T_LC852() {
	nums := []int{0, 1, 0}
	println(peakIndexInMountainArray(nums) == 1)
}
