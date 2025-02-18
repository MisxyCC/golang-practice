package context

import (
	"context"
	"fmt"
	"time"
)

func Demo2() {
	ctx := context.Background()
	doFirst(ctx)
}

func doFirst(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)
	printCh := make(chan int)
	go doSomeStuff(ctx, printCh)
	for num := 1; num <= 3; num++ {
		printCh <- num
	}
	cancelCtx()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("DoFirst: finished\n")

}

func doSomeStuff(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doSomeStuff err: %s\n", err)
			}
			fmt.Printf("doSomeStuff: finished\n")
			return
		case num := <-printCh:
			fmt.Printf("doSomeStuff: value is %d\n", num)
		}
	}
}