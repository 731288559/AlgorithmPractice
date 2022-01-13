package hard

import "sort"

type Flower struct {
	PlantTime int
	GrowTime  int
}

func earliestFullBloom(plantTime []int, growTime []int) int {
	flowers := make([]Flower, 0, len(plantTime))
	for i := range plantTime {
		flowers = append(flowers, Flower{PlantTime: plantTime[i], GrowTime: growTime[i]})
	}
	sort.Slice(flowers, func(i, j int) bool {
		if flowers[i].GrowTime > flowers[j].GrowTime {
			return true
		}
		return false
	})

	var maxTime int
	var plantTimeCur int
	for _, f := range flowers {
		plantTimeCur += f.PlantTime
		if maxTime < plantTimeCur+f.GrowTime {
			maxTime = plantTimeCur + f.GrowTime
		}
		// println(f.GrowTime, f.PlantTime)
	}

	return maxTime
}

func T_LC2136() {
	println(earliestFullBloom([]int{1, 4, 3}, []int{2, 3, 1}), 9)
	println(earliestFullBloom([]int{1, 2, 3, 2}, []int{2, 1, 2, 1}), 9)
	println(earliestFullBloom([]int{1}, []int{1}), 2)
}
