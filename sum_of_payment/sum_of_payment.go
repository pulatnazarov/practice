package sum_of_payment

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SumOfPayment10() {
	var n int
	in, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	in = in[:len(in)-1]
	n, _ = strconv.Atoi(in)

	questions := make(map[int][]int, 0)

	for i := 0; i < n; i++ {
		prices := make([]int, 0)
		number := 0
		in, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		in = in[:len(in)-1]
		number, _ = strconv.Atoi(in) // 12
		number++

		input, err := bufio.NewReader(os.Stdin).ReadString('\n') // 2 2 2 2 2 2 2 3 3 3 3 3
		if err != nil {
			fmt.Println("err is", err)
		}
		input = input[:len(input)-1]

		for _, v := range strings.Split(input, " ") {
			price, _ := strconv.Atoi(v)
			prices = append(prices, price)
		}
		questions[i] = prices
	}

	for _, v := range questions {
		number := 0
		maps := make(map[int]int, len(v))
		for _, item := range v {
			_, exits := maps[item]
			if exits {
				maps[item]++
			} else {
				maps[item] = 1
			}
		}
		result := 0
		for key, val := range maps {
			number = val / 3
			number2 := val % 3
			if number < 1 {
				val *= key
				result += val
			} else {
				a := number*2 + number2
				a *= key
				result += a
			}
		}
		fmt.Println(result)
	}
}

func SumOfPayment5() {
	var n int
	//in, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//in = in[:len(in)-1]
	//n, _ = strconv.Atoi(in)
	fmt.Scan(&n)

	answers := make([]int, 0)
	for i := 0; i < n; i++ {
		var intA, intB = 0, 0
		//in, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		//in = in[:len(in)-1]

		fmt.Scan(&intA, &intB)
		//for i, v := range strings.Split(in, " ") {
		//	if i == 0 {
		//		intA, _ = strconv.Atoi(v)
		//	} else {
		//		intB, _ = strconv.Atoi(v)
		//	}
		//}
		result := intA + intB
		answers = append(answers, result)
	}

	for i := 0; i < len(answers); i++ {
		fmt.Println(answers[i])
	}
}
