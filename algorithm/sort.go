package algorithm

import "fmt"

func quickSort(arr []int, left, right int) {
	if left < right {
		pivot := arr[left]
		low, high := left, right
		for low < high {
			for low < high && arr[high] >= pivot {
				high--
			}
			arr[low], arr[high] = arr[high], arr[low]
			for low < high && arr[low] <= pivot {
				low++
			}
			arr[low], arr[high] = arr[high], arr[low]
		}
		quickSort(arr, left, low-1)
		quickSort(arr, low+1, right)
	}
}

func quickSortV1(arr []int, low, hight int) {
	// 当 low = hight  时跳出
	if low >= hight {
		return
	}

	left, right := low, hight
	pivot := arr[left] // 为了简单起见，直接取左边的第一个数

	for left < right {
		// 先从右边开始迭代

		// 右边的数如果比pivot大，那么就应该将他放在右边，继续向左滑动，遇到一个比他小的为止
		for left < right && arr[right] >= pivot {
			right--
		}

		// 小数移动到左边
		if left < right {
			arr[left] = arr[right]
		}

		// 左边的数如果比pivot小，那么就应该将他放在左边，继续向右滑动，遇到一个比他大的为止
		for left < right && arr[left] < pivot {
			left++
		}

		// 大数移动到又边
		if left < right {
			arr[right] = arr[left]
		}

		// left == right ,pivot 即是最终位置
		if left <= right {
			arr[left] = pivot
		}
	}
	//因为 pivot 的最终位置已锁定

	// 继续排序左边部分
	quickSortV1(arr, low, right-1)
	// 继续排序右边部分
	quickSortV1(arr, right+1, hight)
}

func T_QuickSort() {
	arr := []int{1, 3, 4, 6, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
