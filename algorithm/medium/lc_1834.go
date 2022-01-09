package medium

import (
	"container/heap"
	"fmt"
	"sort"
)

// 1834. 单线程 CPU

type TaskCPU struct {
	StartTime   int
	ProcessTime int
	Num         int
}
type TaskSlice [][]int

func (t TaskSlice) Len() int {
	return len(t)
}
func (t TaskSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t TaskSlice) Less(i, j int) bool {
	return t[i][0] < t[j][0]
}

type TaskHeap []TaskCPU

func (h TaskHeap) Len() int {
	return len(h)
}
func (h TaskHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h TaskHeap) Less(i, j int) bool {
	if h[i].ProcessTime != h[j].ProcessTime {
		return h[i].ProcessTime < h[j].ProcessTime
	}
	return h[i].Num < h[j].Num
}

func (h *TaskHeap) Push(x interface{}) {
	*h = append(*h, x.(TaskCPU))
}

func (h *TaskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getOrder(tasks [][]int) []int {
	ans := make([]int, 0, len(tasks))
	for i := range tasks {
		tasks[i] = append(tasks[i], i)
	}
	sort.Sort(TaskSlice(tasks))
	// fmt.Println(tasks)

	h := &TaskHeap{}
	i := 0
	// TODO: 这里可以优化，避免两段重复代码
	for i < len(tasks) && tasks[i][0] == tasks[0][0] {
		heap.Push(h, TaskCPU{StartTime: tasks[i][0], ProcessTime: tasks[i][1], Num: tasks[i][2]})
		i++
	}
	curTime := 0
	for h.Len() > 0 {
		ts := heap.Pop(h).(TaskCPU)
		ans = append(ans, ts.Num)
		curTime = maxInt(ts.StartTime+ts.ProcessTime, curTime+ts.ProcessTime)
		for i < len(tasks) && tasks[i][0] <= curTime {
			heap.Push(h, TaskCPU{StartTime: tasks[i][0], ProcessTime: tasks[i][1], Num: tasks[i][2]})
			i++
		}
		if h.Len() == 0 && i < len(tasks) {
			tmp := i
			for j := tmp; j < len(tasks) && tasks[j][0] == tasks[tmp][0]; j++ {
				heap.Push(h, TaskCPU{StartTime: tasks[j][0], ProcessTime: tasks[j][1], Num: tasks[j][2]})
				i = j + 1
			}
		}
	}
	return ans
}

func T_LC1834() {
	// fmt.Println(getOrder([][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}), []int{0, 2, 3, 1})
	// fmt.Println(getOrder([][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}), []int{4, 3, 2, 0, 1})

	fmt.Println(getOrder([][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}}), []int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7})
	// fmt.Println(getOrder([][]int{{1, 2}, {5, 2}, {5, 1}, {10, 1}}), []int{0, 2, 1, 3})
}
