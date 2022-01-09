package easy

import (
	"testing"
)

func TestLC852(t *testing.T) {
	nums := []int{0, 1, 0}
	if ans := peakIndexInMountainArray(nums); ans != 1 {
		t.Error()
	}
	nums = []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	if ans := peakIndexInMountainArray(nums); ans != 2 {
		t.Error()
	}
	nums = []int{0, 1, 2, 3}
	if ans := peakIndexInMountainArray(nums); ans != 3 {
		t.Error()
	}
}
