package easy

import "strconv"

// 9. 回文数

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	n := 0
	for x > n {
		n = n*10 + x%10
		x = x / 10
	}
	return x == n || x == n/10
}

func TestLC9() {
	println(isPalindrome(121) == true)
	println(isPalindrome(-121) == false)
	println(isPalindrome(10) == false)
	println(isPalindrome(101) == true)

	println(isPalindrome2(101) == true)
	println(isPalindrome2(1001) == true)
	println(isPalindrome2(1021) == true)
}
