package defer_check

import "fmt"

func Defer() {
	defer func() {
		defer func() {
			fmt.Println("function inside of 1")
		}()
		fmt.Println("function 1")
	}()

	defer func() {
		defer func() {
			fmt.Println("function inside of 2")
		}()
		fmt.Println("function 2")
	}()

	defer func() {
		fmt.Println("function 3")
	}()
}
