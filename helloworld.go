package main

import (
	"fmt"
	"sync"
)

var m sync.Map
var lock sync.Mutex

// 从nil的channel读写会导致panic，同map一样，需要先make初始化
func nilChannel() {
	var ch chan int
	<-ch
	ch <- 1
}

func Test() {
	l := []string{"1", "2"}
	for i := range l {
		l[i] = "3"
	}
	l = append(l, "3")

	fmt.Println(l)

	str := "12345"
	for idx, i := range str {
		fmt.Printf("idx:%d, char:%v\n", idx, i)
	}
}
