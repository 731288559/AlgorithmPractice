package medium

import "fmt"

// 31. 下一个排列

// 思路是仿造一般思维，从低位向高位找到非升序的位A，再从A向右找到比A大中最小的位B，交换AB，最后倒序排A后面的数字
// 例如：538721，先找到3，在找到7，交换37，最后倒置8321，得到571238

func nextPermutation(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}

	var (
		nextExist bool
		temp      int
	)
	for i := length - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			var swapFlag bool
			temp = nums[i]
			for j := i + 1; j < length; j++ {
				if nums[j] <= temp {
					nums[i] = nums[j-1]
					nums[j-1] = temp
					swapFlag = true
					break
				}
			}
			if !swapFlag {
				nums[i] = nums[length-1]
				nums[length-1] = temp
			}

			reverse(nums[i+1:])
			nextExist = true
			break
		}
	}

	if !nextExist {
		reverse(nums)
	}
}

func reverse(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

func TestLC31() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{3, 2, 1}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1, 1, 5}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{1}
	nextPermutation(nums)
	fmt.Println(nums)

	nums = []int{2, 5, 1}
	nextPermutation(nums)
	fmt.Println(nums) // 5 1 2

	nums = []int{1, 2, 4, 3, 1}
	nextPermutation(nums)
	fmt.Println(nums) // 1 3 1 2 4
}
