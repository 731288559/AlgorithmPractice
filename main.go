package main

import (
	"strconv"

	"my_practice/algorithm/easy"
	"my_practice/algorithm/hard"
	"my_practice/algorithm/medium"
)

func t() {
	easy.T_LC66()
	medium.T_LC230()
	hard.T_LC282()
	num := "123"
	println(strconv.Atoi(num[0:1]))
}

func main() {
	// easy.T_LC5976()
	// feature.DeferTest()
	// hard.T_LC1036()
	// medium.T_LC926()
	// algorithm.T_QuickSort()

	// feature.SliceTest1()
	// feature.Channel_3()
	// medium.T_LC334()

	easy.T_LC747()
}
