package medium

import (
	"strconv"
)

// 299. 猜数字游戏

func getHint(secret string, guess string) string {
	a, b := 0, 0
	l1 := make([]int, 10)
	l2 := make([]int, 10)
	for i := range guess {
		if guess[i] == secret[i] {
			a++
		} else {
			l1[guess[i]-'0']++
			l2[secret[i]-'0']++
		}
	}
	for i := range l1 {
		if l1[i] <= l2[i] {
			b += l1[i]
		} else {
			b += l2[i]
		}
	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}

func getHint_v2(secret string, guess string) string {
	a, b := 0, 0
	visit := make([]bool, len(secret))
	for i := range guess {
		if guess[i] == secret[i] {
			if !visit[i] {
				// println("i, a, b = ", i, a, b)
				a++
				visit[i] = true
			}
		} else {
			for j := range secret {
				if guess[i] == secret[j] {
					if secret[j] == guess[j] {
						if !visit[j] {
							// println("a, b, i, j = ", a, b, i, j)
							a++
							visit[j] = true
						}
					} else {
						if !visit[j] {
							b++
							visit[j] = true
							break
						}
					}
				}
			}
		}
		// fmt.Println("a, b, visit = ", a, b, visit)
	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}

func T_LC299() {
	println(getHint("1807", "7810"), "1A3B")
	println(getHint("1123", "0111"), "1A1B")
	println(getHint("1", "0"), "0A0B")
	println(getHint("1", "1"), "1A0B")
}
