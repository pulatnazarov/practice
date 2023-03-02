package concurrency

import (
	"fmt"
	"time"
)

func Channel() {
	readChan := make(chan int)
	writtenChan := make(chan int)

	go func() {
		for {

		}
	}()

	for {
		select {
		case v := <-readChan:
			fmt.Println("case read", v)
		case writtenChan <- 10:
			fmt.Println("case written")
		case <-time.After(time.Second * 5):
			fmt.Println("goodbye")
			return
		}
	}
}
