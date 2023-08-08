package leetcode

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	maxCost := 0
	tempInformTime := make([]int, n)
	for i, n := range manager {
		_ = i
		tempCost := 0
		tempLeader := n
		for tempLeader != -1 {
			tempCost += informTime[tempLeader]
			tempLeader = manager[tempLeader]
		}
		tempInformTime[i] = tempCost
		if tempCost > maxCost {
			maxCost = tempCost
		}
	}
	return maxCost
}
