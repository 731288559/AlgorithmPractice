package easy

func maxDepth1614(s string) int {
	var res int
	var tmp int
	for _, v := range s {
		if v == '(' {
			tmp++
			if tmp > res {
				res = tmp
			}
		} else if v == ')' {
			tmp--
		}
	}

	return res
}

func T_LC1614() {
	println(maxDepth1614("(1+(2*3)+((8)/4))+1"), 3)
	println(maxDepth1614("(1)+((2))+(((3)))"), 3)
	println(maxDepth1614("1+(2*3)/(2-1)"), 1)
	println(maxDepth1614("1"), 0)
}
