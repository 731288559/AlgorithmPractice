package hard

import "container/heap"

// 407. 接雨水 II

func trapRainWater(heightMap [][]int) int {
	row := len(heightMap)
	col := len(heightMap[0])
	if row < 3 || col < 3 {
		return 0
	}
	h := &hp{}
	visit := make([][]bool, row)
	for i := range visit {
		visit[i] = make([]bool, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 || i == row-1 || j == 0 || j == col-1 {
				heap.Push(h, cell{i, j, heightMap[i][j]})
				visit[i][j] = true
			}
		}
	}
	var result int
	dirs := []int{-1, 0, 1, 0, -1}
	// 利用最小堆，每次处理最矮的围墙，逐渐缩小范围
	for h.Len() > 0 {
		curCell := heap.Pop(h).(cell)
		for k := 0; k < 4; k++ {
			curX, curY := curCell.x+dirs[k], curCell.y+dirs[k+1]
			if 0 <= curX && curX < row && 0 <= curY && curY < col && !visit[curX][curY] {
				if heightMap[curX][curY] < curCell.z {
					result += curCell.z - heightMap[curX][curY]
				}
				visit[curX][curY] = true
				heap.Push(h, cell{curX, curY, maxInt(curCell.z, heightMap[curX][curY])})
			}
		}
	}
	return result
}

type cell struct{ x, y, z int }
type hp []cell

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].z < h[j].z }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(cell)) }
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func T_LC407() {
	println(trapRainWater([][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}), 4)
	println(trapRainWater([][]int{
		{3, 3, 3, 3, 3},
		{3, 2, 2, 2, 3},
		{3, 2, 1, 2, 3},
		{3, 2, 2, 2, 3},
		{3, 3, 3, 3, 3},
	}), 10)
	println(trapRainWater([][]int{
		{12, 13, 1, 12},
		{13, 4, 13, 12},
		{13, 8, 10, 12},
		{12, 13, 12, 12},
		{13, 13, 13, 13},
	}), 14)
}
