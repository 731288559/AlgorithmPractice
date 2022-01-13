package easy

// 5930. 两栋颜色不同且距离最远的房子

func maxDistance(colors []int) int {
	a := colors[0]
	b := colors[len(colors)-1]

	lenA, lenB := len(colors)-1, len(colors)-1
	for i := 0; i < len(colors); i++ {
		if colors[i] != b {
			break
		}
		lenB--
	}
	for i := len(colors) - 1; i > 0; i-- {
		if colors[i] != a {
			break
		}
		lenA--
	}
	if lenA > lenB {
		return lenA
	}
	return lenB
}

func T_LC5930() {
	println(maxDistance([]int{1, 1, 1, 6, 1, 1, 1}), 3)
	println(maxDistance([]int{1, 8, 3, 8, 3}), 4)
	println(maxDistance([]int{0, 1}), 1)
	println(maxDistance([]int{62, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 62}), 11)
	println(maxDistance([]int{4, 4, 4, 11, 4, 4, 11, 4, 4, 4, 4, 4}), 8)
}
