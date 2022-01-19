package easy

func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]struct{})
	for i, n := range nums {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[n]; ok {
			return true
		}
		set[n] = struct{}{}
	}
	return false
}

func T_LC219() {
	println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3), true)
	println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1), true)
	println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2), false)
}
