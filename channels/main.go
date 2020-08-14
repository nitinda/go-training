package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	// Channel with buffered space
	chInt := make(chan int, 50)

	// for j := 0; j < 5; j++ {
	wg.Add(2)

	// Data flow from channel
	go func(chInt <-chan int) {
		// For range loop
		for j := range chInt {
			fmt.Println(j)
		}
		wg.Done()
	}(chInt)

	go func(chInt chan<- int) {
		chInt <- 32
		chInt <- 33
		close(chInt)
		wg.Done()
	}(chInt)
	// }

	wg.Wait()

	go logger()
	// Defer function call
	// defer func() {
	// 	close(logCh)
	// }()

	logCh <- logEntry{time.Now(), logInfo, "App is starting....."}
	logCh <- logEntry{time.Now(), logInfo, "App is shuting down....."}

	doneCh <- struct{}{}
}
