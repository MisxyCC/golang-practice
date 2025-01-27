package concurrency

import (
	"fmt"
	"sync"
)

func generatingNumber(total int, ch chan<- int) {
	
	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Sending the data %d to the channel\n", idx)
		ch <- idx
	}
}

func printNumber(idx int, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("%d: Read %d from channel\n", idx, num)
	}
}

func ExecuteSampleGoroutinesAndChannels() {
	numberChan := make(chan int)
	var wg sync.WaitGroup
	
	for idx := 1; idx <= 3; idx++ {
		wg.Add(1)
		go printNumber(idx, &wg, numberChan)
	}
	generatingNumber(5, numberChan)
	close(numberChan)
	

	fmt.Println("Waiting the goroutines to finish their tasks")
	wg.Wait()
	fmt.Println("The goroutines finished their tasks, Done!")
}