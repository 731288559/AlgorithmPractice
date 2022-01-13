package other

// 给定一个数组nums,nums[i]表示每个木头的长度,木头可以截断,现在需要k根长度一样的木头,每根木头的最大长度是多少?

func kWood(nums []int, k int) int {
	min := 10000
	max := 0
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		}
	}
	// println(min, max)

	m := 0
	count := 0
	for min < max {
		m = (min + max + 1) / 2
		count = 0
		for i := range nums {
			count += nums[i] / m
		}
		if count >= k {
			min = m
		} else {
			max = m - 1
		}
		// println(min, max)
	}

	return min
}

func TestOther1() {
	nums := []int{4, 7, 2, 10, 5}
	println(kWood(nums, 5)) // 4
}
