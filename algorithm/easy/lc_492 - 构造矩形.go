package easy

import (
	"fmt"
	"math"
)

// 492. 构造矩形

func constructRectangle(area int) []int {
	for i := int(math.Sqrt(float64(area))); i > 0; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}
	return nil
}

func T_LC492() {
	fmt.Println(constructRectangle(4))
	fmt.Println(constructRectangle(9))
	fmt.Println(constructRectangle(10))
}
