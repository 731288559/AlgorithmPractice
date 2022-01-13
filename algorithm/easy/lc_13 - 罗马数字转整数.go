package easy

import "fmt"

// 13. 罗马数字转整数

func romanToInt(s string) int {
	var result int
	last := 1001
	for i := range s {
		n := getNum(s[i])
		if n <= last {
			result += n
			last = n
		} else {
			result = result - 2*last
			result += n
		}
		last = n
	}
	fmt.Println("input, output", s, result)
	return result
}

func getNum(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func TestLC13() {
	println(romanToInt("III") == 3)
	println(romanToInt("IV") == 4)
	println(romanToInt("IX") == 9)
	println(romanToInt("LVIII") == 58)
	println(romanToInt("MCMXCIV") == 1994)
}
