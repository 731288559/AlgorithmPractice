package feature

import (
	"fmt"
	"sync"
	"time"
)

// ChannelTest1 两个goroutine交替打印1~20
func ChannelTest1() {
	n := 20
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 1; i <= n; i = i + 2 {
			ch1 <- i
			fmt.Println("goruntine 1 print:", <-ch2)
			// if n, ok := <-ch2; ok {
			// 	fmt.Println("channel 1 print:", n)
			// }
		}
		wg.Done()
	}()

	go func() {
		for i := 2; i <= n; i = i + 2 {
			fmt.Println("goruntine 2 print:", <-ch1)
			ch2 <- i
		}
		wg.Done()
	}()
	wg.Wait()
}

// ChannelTest2 两个协程交替打印0和1，共10次
func ChannelTest2() {
	ch := make(chan int, 0)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- 0
			fmt.Println("goruntine 1 print:", <-ch)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			fmt.Println("goruntine 2 print:", <-ch)
			ch <- 1
		}
		wg.Done()
	}()

	wg.Wait()
}

func Channel_3() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		select {
		case x, ok := <-c:
			if ok {
				println(x)
			}
		}
	}
}
