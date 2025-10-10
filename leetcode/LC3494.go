package leetcode

func minTime(skill []int, mana []int) int64 {
	n := len(skill)
	lastFinish := make([]int, n+1) // 前n个巫师完成药水的时间

	for _, m := range mana {
		sumTemp := 0
		for i, s := range skill {
			sumTemp = max(sumTemp, lastFinish[i+1]) + s*m
		}
		lastFinish[n] = sumTemp

		for i := n - 1; i > 0; i-- {
			lastFinish[i] = lastFinish[i+1] - skill[i]*m
		}
	}

	return int64(lastFinish[n])
}
