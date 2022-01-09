package easy

import "fmt"

// 496. 下一个更大元素 I

func nextGreaterElement_v1(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1))
	for i := range nums1 {
		n := -1
		startFlag := false
		for j := range nums2 {
			if nums2[j] == nums1[i] {
				startFlag = true
				continue
			}
			if startFlag {
				if nums2[j] > nums1[i] {
					n = nums2[j]
					break
				}
			}
		}
		result = append(result, n)
	}
	return result
}

func nextGreaterElement(nums1, nums2 []int) []int {
	result := make([]int, len(nums1), len(nums1))

	m := make(map[int]int)
	stack := make([]int, 0, len(nums2))
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			m[nums2[i]] = stack[len(stack)-1]
		} else {
			m[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])
	}
	for i, n := range nums1 {
		result[i] = m[n]
	}
	return result
}

func T_LC496() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
