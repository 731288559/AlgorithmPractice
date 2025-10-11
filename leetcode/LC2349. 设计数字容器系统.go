package leetcode

import (
	"container/heap"
	"sort"
	"strings"
)

// 2349. 设计数字容器系统

type NumberContainers struct {
	NumberToIndexMap map[int]*hp
	IndexToNumberMap map[int]int
}

func Constructor() NumberContainers {
	return NumberContainers{
		NumberToIndexMap: make(map[int]*hp),
		IndexToNumberMap: make(map[int]int),
	}
}

func (this *NumberContainers) Change(index int, number int) {
	_, ok := this.NumberToIndexMap[number]
	if !ok {
		this.NumberToIndexMap[number] = &hp{}
	}
	heap.Push(this.NumberToIndexMap[number], index)
	this.IndexToNumberMap[index] = number
}

func (this *NumberContainers) Find(number int) int {
	indexList, ok := this.NumberToIndexMap[number]
	if !ok {
		return -1
	}

	for indexList.Len() > 0 && this.IndexToNumberMap[indexList.IntSlice[0]] != number {
		heap.Pop(indexList)
	}

	if indexList.Len() == 0 {
		return -1
	}
	return indexList.IntSlice[0]
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v any) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }

func canBeTypedWords(text string, brokenLetters string) int {
	cnt := 0
	for _, s := range strings.Split(text, " ") {
		if !strings.ContainsAny(s, brokenLetters) {
			cnt++
		}
	}
	return cnt
}

func maxFreqSum(s string) int {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}

	max1, max2 := 0, 0
	for k, v := range m {
		switch k {
		case 'a', 'e', 'i', 'o', 'u':
			max1 = max(max1, v)
		default:
			max2 = max(max2, v)
		}
	}
	return max1 + max2
}
