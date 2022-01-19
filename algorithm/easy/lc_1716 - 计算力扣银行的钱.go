package easy

func totalMoney2(n int) int {
	var res int
	week := 0
	for i := 0; i < n; i++ {
		if i%7 == 0 {
			week++
		}
		res += i%7 + week

	}
	return res
}

func totalMoney(n int) int {
	week := n / 7
	firstWeekMoney := (1 + 7) * 7 / 2
	lastWeekMoney := firstWeekMoney + 7*(week-1)

	day := n % 7
	firstDayMoney := 1 + week
	lastDayMoney := firstDayMoney + day - 1
	return (firstWeekMoney+lastWeekMoney)*week/2 + (firstDayMoney+lastDayMoney)*day/2
}

func T_LC1716() {
	println(totalMoney(4), 10)
	println(totalMoney(10), 37)
	println(totalMoney(20), 96)
}
