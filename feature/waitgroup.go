package feature

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

func WaitGroupTest1() {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		wg.Done()
	}
	wg.Wait()
	println("main exit")
}

// WaitGroupTest2 使用m个协程完成n个任务
func WaitGroupTest2() {
	producer := func(n int, taskCh chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		defer close(taskCh)

		for i := range n {
			taskCh <- i
			fmt.Println("produce task:", i)
		}
		time.Sleep(time.Second)
	}

	consumer := func(i int, taskCh <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		for task := range taskCh {
			time.Sleep(time.Millisecond * 300)
			fmt.Printf("consumer %d finish task %d\n", i, task)
		}

	}

	m := 3
	n := 10

	taskCh := make(chan int, n)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go producer(n, taskCh, wg)

	wg.Add(m)
	for i := range m {
		go consumer(i, taskCh, wg)
	}

	wg.Wait()
	fmt.Println("all task done!")
}

// WaitGroupTest3 使用m个协程完成n个任务，错误处理版本
func WaitGroupTest3() {
	producer := func(ctx context.Context, n int, taskCh chan<- int, wg *sync.WaitGroup, errCh chan<- error) {
		defer wg.Done()
		defer close(taskCh)

		for i := range n {
			select {
			case <-ctx.Done():
				return
			default:
				taskCh <- i
				fmt.Println("produce task:", i)
			}
		}

	}

	consumer := func(ctx context.Context, i int, taskCh <-chan int, wg *sync.WaitGroup, errCh chan<- error) {
		defer wg.Done()

		for task := range taskCh {
			fmt.Printf("consumer %d get task %d\n", i, task)
			select {
			case <-ctx.Done():
				fmt.Printf("receive cancel signal, exit consumer: %d drop task %d\n", i, task)
				return
			default:
				time.Sleep(time.Millisecond * 200)
				var err error
				if task == 5 {
					err = errors.New("error mock")
				}
				if err != nil {
					select {
					case errCh <- err:
						//fmt.Printf("task:%d error:%s\n", task, err.Error())
					default: // 防止阻塞
					}
					fmt.Println("error occur, exit consumer:", i)
					return
				}
				fmt.Printf("consumer %d finish task %d\n", i, task)
			}

		}

	}

	m := 3
	n := 10

	taskCh := make(chan int, n)
	errCh := make(chan error, 1)
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go producer(ctx, n, taskCh, wg, errCh)
	time.Sleep(time.Millisecond * 200)

	wg.Add(m)
	for i := range m {
		go consumer(ctx, i, taskCh, wg, errCh)
	}

	//wg.Wait()

	select {
	case err := <-errCh:
		fmt.Printf("catch err: %v，cancel all task...\n", err)
		cancel()
	case <-waitGroupDone(wg):
		fmt.Println("all task success!")
	}

	wg.Wait()

}

var group singleflight.Group

// 辅助函数：将 WaitGroup 转换为 Channel
func waitGroupDone(wg *sync.WaitGroup) <-chan struct{} {
	ch := make(chan struct{})
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}
