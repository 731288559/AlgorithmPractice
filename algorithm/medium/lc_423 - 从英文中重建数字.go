package medium

import "strconv"

// 423. 从英文中重建数字

func originalDigits(s string) string {
	var ans string
	// _ := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	var cnt [26]int
	for i := range s {
		cnt[s[i]-'a']++
	}

	var nums [10]int
	nums[0] = cnt['z'-'a']
	nums[6] = cnt['x'-'a']
	nums[7] = cnt['s'-'a'] - nums[6]
	nums[2] = cnt['w'-'a']
	nums[5] = cnt['v'-'a'] - nums[7]
	nums[8] = cnt['g'-'a']
	nums[9] = cnt['i'-'a'] - nums[5] - nums[6] - nums[8]
	nums[1] = cnt['n'-'a'] - nums[7] - 2*nums[9]
	nums[4] = cnt['r'-'a'] - nums[0]
	nums[3] = cnt['h'-'a'] - nums[8]

	for i := 0; i < 10; {
		if nums[i] > 0 {
			ans += strconv.Itoa(i)
			nums[i]--
		} else {
			i++
		}
	}
	return ans
}

func T_LC423() {
	// println(originalDigits("owoztneoer"), "012")
	println(originalDigits("fviefuro"), "45")
	println(originalDigits("zerozero"), "00")
}
