package palindrome

import (
	"fmt"
	"math"
	"strconv"
)

func Palindrome() {
	a := 10
	l := len(strconv.Itoa(a))
	left := a / int(math.Pow10(l/2))
	right := a % int(math.Pow10(l/2))
	if l%2 != 0 {
		left /= 10
	}
	if len(strconv.Itoa(right)) != l/2 {
		right += int(math.Pow10(l / 2))
		left = left*10 + 1
	}

	lr := len(strconv.Itoa(right))
	for i := 0; i < l/2; i++ {
		if left%10 == right/int(math.Pow10(lr-1)) {
			left /= 10
			right %= int(math.Pow10(lr - 1))
			lr--
		} else {
			fmt.Println("not palindrome")
			return
		}
	}
	fmt.Println("palindrome")
}
