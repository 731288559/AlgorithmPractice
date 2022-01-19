package medium

func nthUglyNumber(n int) int {
	l := make([]int, n+1)
	l[1] = 1

	i2, i3, i5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		a, b, c := l[i2]*2, l[i3]*3, l[i5]*5
		min := min(min(a, b), c)
		if min == a {
			i2++
		}
		if min == b {
			i3++
		}
		if min == c {
			i5++
		}
		l[i] = min
	}
	return l[n]
}

func T_LC264() {
	println(nthUglyNumber(10), 12)
	println(nthUglyNumber(1), 1)
}
