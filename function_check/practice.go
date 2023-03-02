package function_check

import "fmt"

func Function() {
	f := func() {
		fmt.Println("prints string")
	}

	f()
}
