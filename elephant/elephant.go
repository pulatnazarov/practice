package elephant

import (
	"fmt"
	"time"
)

func Elephant() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(f1())
	fmt.Println(time.Now())

	fmt.Println("***")
	t = time.Now()
	fmt.Println(t)
	fmt.Println(f2())
	fmt.Println(time.Now())
}

func f2() int {
	x, y, count := 0, 5, 0

	fmt.Scan(&x)

	for x != 0 {
		if x-y >= 0 {
			count++
			x -= y
		} else {
			y--
		}
	}

	return count
}

func f1() int {
	var (
		count = 0
		x     = 0
		y     = 5
		q     = 5
		b     = 0
	)
	fmt.Scan(&x)
	for x != 0 {
		if q < y {
			y--
		} else {
			q = x % y
			b = x / y
			count += b
			x = q
			if x == 0 {
				break
			}
		}
	}

	return count
}
