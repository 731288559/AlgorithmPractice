package easy

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}

	res := make([][]int, 0, m)
	var tmp = make([]int, 0, n)
	for i := range original {
		tmp = append(tmp, original[i])
		if i%n == n-1 {
			res = append(res, tmp)
			tmp = make([]int, 0, n)
		}
	}
	return res
}

func T_LC2022() {
	fmt.Println(construct2DArray([]int{1, 2, 3, 4}, 2, 2))
	fmt.Println(construct2DArray([]int{1, 2, 3}, 1, 3))
	fmt.Println(construct2DArray([]int{1, 2}, 1, 1))
	fmt.Println(construct2DArray([]int{3}, 1, 2))
}
