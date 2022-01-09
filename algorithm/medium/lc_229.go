package medium

import "fmt"

// 229. 求众数 II

func majorityElement(nums []int) []int {
	var (
		a  int
		b  int
		c1 int
		c2 int
	)

	for _, i := range nums {
		if c1 > 0 && a == i {
			c1++
		} else if c2 > 0 && b == i {
			c2++
		} else if c1 == 0 {
			a = i
			c1++
		} else if c2 == 0 {
			b = i
			c2++
		} else {
			c1--
			c2--
		}
	}

	c1 = 0
	c2 = 0
	for _, i := range nums {
		if i == a {
			c1++
		} else if i == b {
			c2++
		}
	}
	var result []int
	if c1 > len(nums)/3 {
		result = append(result, a)
	}
	if c2 > len(nums)/3 {
		result = append(result, b)
	}
	return result
}

func T_LC229() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
	nums = []int{1}
	fmt.Println(majorityElement(nums))
	nums = []int{1, 1, 1, 3, 3, 2, 2, 2}
	fmt.Println(majorityElement(nums))
}
