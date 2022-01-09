package easy

import "fmt"

func slowestKey(releaseTimes []int, keysPressed string) byte {
	var res byte

	var p int
	var maxTime int
	for i := range releaseTimes {
		if releaseTimes[i]-p > maxTime || (releaseTimes[i]-p == maxTime && keysPressed[i] > res) {
			maxTime = releaseTimes[i] - p
			res = keysPressed[i]
		}
		p = releaseTimes[i]
	}

	return res
}

func T_LC1629() {
	fmt.Println(string(slowestKey([]int{9, 29, 49, 50}, "cbcd")))
	fmt.Println(string(slowestKey([]int{12, 23, 36, 46, 62}, "spuda")))
}
