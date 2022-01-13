package medium

// 649. Dota2 参议院

// 考虑 DR DR DR DR RR 的情况，不能直接判断R个数>len/2+1直接判断R胜利

func predictPartyVictory(senate string) string {
	l1 := make([]int, 0, len(senate))
	l2 := make([]int, 0, len(senate))
	for i, n := range senate {
		if n == 'R' {
			l1 = append(l1, i)
		} else {
			l2 = append(l2, i)
		}
	}

	for len(l1) > 0 && len(l2) > 0 {
		if l1[0] < l2[0] {
			l1 = append(l1, l1[0]+len(senate))
		} else {
			l2 = append(l2, l2[0]+len(senate))
		}
		l1 = l1[1:]
		l2 = l2[1:]
	}

	if len(l1) > 0 {
		return "Radiant"
	}
	return "Dire"
}

func T_LC649() {
	println(predictPartyVictory("RD"), "Radiant")
	println(predictPartyVictory("RDD"), "Dire")
}
