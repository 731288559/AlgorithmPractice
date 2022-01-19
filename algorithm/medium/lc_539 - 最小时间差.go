package medium

import (
	"math"
	"sort"
	"strconv"
)

func findMinDifference(timePoints []string) int {
	var min int = math.MaxInt
	sort.Strings(timePoints)

	if len(timePoints) < 2 {
		return -1
	}
	for i := 1; i < len(timePoints); i++ {
		d := getTimeInterval(timePoints[i], timePoints[i-1])
		if d < min {
			min = d
		}
	}
	d := getTimeInterval(timePoints[0], "00:00") + getTimeInterval("24;00", timePoints[len(timePoints)-1])
	if d < min {
		min = d
	}
	return min
}

func getTimeInterval(a, b string) int {
	ha, _ := strconv.Atoi(a[:2])
	ma, _ := strconv.Atoi(a[3:])

	hb, _ := strconv.Atoi(b[:2])
	mb, _ := strconv.Atoi(b[3:])

	return ha*60 + ma - hb*60 - mb
}

func T_LC539() {
	println(findMinDifference([]string{"23:59", "00:00"}), 1)
	println(findMinDifference([]string{"00:00", "23:59", "00:00"}), 0)
}
