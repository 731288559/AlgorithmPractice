package easy

// 859. 亲密字符串

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) || len(s) < 2 {
		return false
	}
	var cnt, cnt2 [26]int
	var f bool
	var diff int
	for i := range s {
		if s[i] != goal[i] {
			diff++
		}
		cnt[s[i]-'a']++
		cnt2[goal[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if cnt[i] != cnt2[i] {
			return false
		}
		if cnt[i] >= 2 {
			f = true
		}
	}

	return diff == 2 || diff == 0 && f
}

func T_LC859() {
	println(buddyStrings("ab", "ba"), true)
	println(buddyStrings("ab", "ab"), false)
	println(buddyStrings("aa", "aa"), true)
	println(buddyStrings("aaaaaaabc", "aaaaaaacb"), true)
}
