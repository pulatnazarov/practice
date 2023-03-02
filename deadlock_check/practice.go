package deadlock_check

import (
	"fmt"
	"math/rand"
	"time"
)

func Deadlock() {
	words := []string{
		"hello", "apple", "a", "a", "hello", "word", "www", "www",
	}

	for word := range removeRepeated(generator(words)) {
		fmt.Println("word: ", word)
	}

	//nums := sendRandomNumbers()
	//for num := range nums {
	//	fmt.Println(num)
	//}
}

func generator(words []string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, word := range words {
			ch <- word
		}
	}()

	return ch
}

func removeRepeated(inputStream <-chan string) <-chan string {
	outputStream := make(chan string)
	go func() {
		defer close(outputStream)
		temp := ""

		for val := range inputStream {
			if val != temp {
				outputStream <- val
			}
			temp = val
		}
	}()

	return outputStream
}

func sendRandomNumbers() <-chan int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10) + 1
	nums := make(chan int)
	go func() {
		defer close(nums)
		for i := 0; i < n; i++ {
			nums <- rand.Intn(50)
		}
	}()

	return nums
}
