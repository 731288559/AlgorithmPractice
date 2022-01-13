package easy

// 495. 提莫攻击

func findPoisonedDuration(timeSeries []int, duration int) int {
	// if len(timeSeries) == 0 {
	// 	return 0
	// }
	// d := duration
	// maxT := timeSeries[0] + duration - 1
	// for _, n := range timeSeries {
	// 	if n > maxT {
	// 		d += duration
	// 	} else {
	// 		d += duration + n - maxT - 1
	// 	}
	//  maxT = n + duration - 1
	// }
	// return d

	d := 0
	expired := 0
	for _, n := range timeSeries {
		if n >= expired {
			d += duration
		} else {
			d += n + duration - expired
		}
		expired = n + duration
	}
	return d
}

func T_LC495() {
	// println(findPoisonedDuration([]int{1, 4}, 2), 4)
	println(findPoisonedDuration([]int{1, 2}, 2), 3)
}
