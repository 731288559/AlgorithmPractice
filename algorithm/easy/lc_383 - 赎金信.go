package easy

// 383. 赎金信

func canConstruct(ransomNote string, magazine string) bool {
	arr := [26]int{}
	for i := range ransomNote {
		arr[ransomNote[i]-'a']++
	}
	for i := range magazine {
		if arr[magazine[i]-'a'] > 0 {
			arr[magazine[i]-'a']--
		}
	}
	for i := range arr {
		if arr[i] > 0 {
			return false
		}
	}

	return true
}
