package feature

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func ErrGroupTest() {
	eg, ctx := errgroup.WithContext(context.Background())

	producer := func(ctx context.Context, n int, taskCh chan<- int) {
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

	consumer := func(ctx context.Context, i int, taskCh <-chan int) error {
		for task := range taskCh {
			fmt.Printf("consumer %d get task %d\n", i, task)
			select {
			case <-ctx.Done():
				fmt.Printf("receive cancel signal, exit consumer: %d drop task %d\n", i, task)
				return nil
			default:
				time.Sleep(time.Millisecond * 200)
				var err error
				if task == 5 {
					err = errors.New("error mock")
				}
				if err != nil {
					fmt.Println("error occur, exit consumer:", i)
					return err
				}
				fmt.Printf("consumer %d finish task %d\n", i, task)
			}

		}
		return nil
	}

	m := 3
	n := 10

	taskCh := make(chan int, n)

	go producer(ctx, n, taskCh)
	time.Sleep(time.Millisecond * 200)

	for i := range m {
		eg.Go(
			func() error {
				return consumer(ctx, i, taskCh)
			},
		)
	}

	if err := eg.Wait(); err != nil {
		fmt.Println("err:", err)
	}
}
