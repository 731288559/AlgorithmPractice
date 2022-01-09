package easy

// 476. 数字的补数

func findComplement(num int) int {
	result := 0
	flag := false
	for i := 31; i >= 0; i-- {
		t := num >> i & 1
		if t == 0 && flag == false {
			continue
		}
		if t == 1 {
			flag = true
		} else {
			result += 1 << i
		}
	}
	return result
}

func T_LC476() {
	println(findComplement(5))
	println(findComplement(1))
	println(findComplement(10))
}
