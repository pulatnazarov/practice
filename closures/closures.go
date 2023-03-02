package closures

import "fmt"

func Closures() int {
	var a = "some string"
	var x = 0

	x = 1

	closure := func() int {
		fmt.Println(a)
		x++
		return x
	}

	return closure()
}
