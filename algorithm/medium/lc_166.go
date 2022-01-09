package medium

import "strconv"

// 166. 分数到小数

func fractionToDecimal(numerator int, denominator int) string {
	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}
	var result string
	var flag bool

	if numerator < 0 {
		numerator = -numerator
		flag = !flag
	}
	if denominator < 0 {
		denominator = -denominator
		flag = !flag
	}

	big := numerator / denominator
	small := ""
	left := numerator - big*denominator
	m := make(map[int]int)
	loopStartIdx := -1
	i := 0
	for left != 0 {
		if idx, ok := m[left]; ok {
			loopStartIdx = idx
			break
		} else {
			m[left] = i
		}
		k := left * 10 / denominator
		left = left*10 - k*denominator
		small += strconv.Itoa(k)
		i++
	}
	if loopStartIdx >= 0 {
		small = small[:loopStartIdx] + "(" + small[loopStartIdx:] + ")"
	}

	result = strconv.Itoa(big) + "." + small
	if flag {
		result = "-" + result
	}
	return result
}

func T_LC166() {
	println(fractionToDecimal(1, 2))
	println(fractionToDecimal(2, 1))
	println(fractionToDecimal(2, 3))
	println(fractionToDecimal(4, 333))
	println(fractionToDecimal(1, 5))
	println(fractionToDecimal(2, 15))
	println(fractionToDecimal(1, 6))
}
