package easy

import "fmt"

// 66. 加一

func plusOne(digits []int) []int {
	var f bool
	digits[len(digits)-1]++
	for i := len(digits) - 1; i >= 0; i-- {
		if f {
			digits[i]++
			f = false
		}

		if digits[i] >= 10 {
			digits[i] -= 10
			f = true
		}
	}
	if f {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func T_LC66() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))
	digits = []int{0}
	fmt.Println(plusOne(digits))
	digits = []int{9, 9}
	fmt.Println(plusOne(digits))
}
