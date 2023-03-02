package goroutine

import "fmt"

func GoRoutine() {

	a := []int{6, 7, 8, 9, 10}
	ch := make(chan int)
	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}

	/*
		ch1 := make(chan int)
		ch2 := make(chan int)
		fmt.Println("interesno deadlock bomadi")
		go func() {
			v := 1
			ch1 <- v
			v2 := <-ch2
			fmt.Println(v, v2)
		}()

		v := 2
		ch2 <- v
		v2 := <-ch1
		fmt.Println(v, v2)
	*/
	/*
		ch1 := make(chan int)
		ch2 := make(chan int)
		go func() {
			v := 1
			ch1 <- v
			v2 := <-ch2
			fmt.Println(v, v2)
		}()

		v := 2
		var v2 int
		select {
		case ch2 <- v:
		case v2 = <-ch1:
		}
		fmt.Println(v, v2)
	*/
}
