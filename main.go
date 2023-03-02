package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math"
	"practice/binary_search"
	"practice/client_request"
	"practice/closures"
	"practice/concurrency"
	"practice/cron_job"
	"practice/deadlock_check"
	"practice/defer_check"
	"practice/elephant"
	"practice/function_check"
	"practice/goroutine"
	"practice/linear_search"
	"practice/morse"
	"practice/palindrome"
	"practice/recover_check"
	"practice/recursive"
	"practice/sum_of_payment"
	"practice/type_assertion"
	"practice/type_switch"
)

// Main function
func main() {
	//Recursive()
	//GoRoutine()
	//leetcode.DeleteDuplicates()
	//leetcode.ClimbStairs()
	//leetcode.AddBinary()
	//leetcode.Remove()
	//leetcode.LinkedList()
	//Promotion()
	//Palindrome()
	//Morse()
	//BinarySearch()
	//LinearSearch()
	//SumOfPayment()
	//Elephant()
	//Closures()
	//ClientRequest()
	//Concurrency()
	//Cron()
	//array()
	//mapHabr()
	//pi()
	//DeadLock()
	//TypeSwitch()
	//TypeAssertion()
	//Function()
	//Defer()
	//Recover()
	fmt.Println()
}

func Recursive() {
	f := recursive.Recursive(1)
	//recursive.Print(f)
	for _, v := range f {
		recursive.New(v.Id)
	}
}

func GoRoutine() {
	goroutine.GoRoutine()
}

func Promotion() {
	t := 0
	fmt.Scan(&t)
	answ := []int{}
	for t > 0 {
		t--
		var a, b, n, m int
		sum := 0
		fmt.Scan(&a, &b, &n, &m)
		if m > n {
			answ = append(answ, a*n)
		} else {
			for n > m {
				sum += m * a
				n -= m
				n -= 1
			}
			sum += n * b
		}

		if sum != 0 {
			answ = append(answ, sum)
		}
	}

	for _, v := range answ {
		fmt.Println(v)
	}
}

func Palindrome() {
	palindrome.Palindrome()
}

func Morse() {
	morse.Morse()
}

func BinarySearch() {
	binary_search.BinarySearch()
}

func LinearSearch() {
	linear_search.LinearSearch()
}

func SumOfPayment() {
	//sum_of_payment.SumOfPayment10()
	sum_of_payment.SumOfPayment5()
}

func Elephant() {
	elephant.Elephant()
}

func Closures() {
	fmt.Println(closures.Closures())
}

func ClientRequest() {
	client_request.ClientRequest()
}

func GeneratePasswordHash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), 10)
}

func Concurrency() {
	concurrency.Channel()
}

func Cron() {
	cron_job.RunCronJobs()
}

func mapHabr() {
	var m = map[string]int{}
	for _, word := range []string{"hello", "world", "from", "the",
		"best", "language", "in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		println(k, v)
	}
}

func array() {
	a := make([]int, 0, 10)
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	fmt.Println(a)
	fmt.Printf("%p", a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	b := a[2:5]
	b[0] = 100

	fmt.Println(b)
	fmt.Printf("%p", b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
}

func pi() {
	pi := math.Pi
	fmt.Printf("1 - %v | 2 - %.7g | 3 - %.5f | 4 - (%6.2f) | 5 - %e\n", pi, pi, pi, pi, pi)
}

func DeadLock() {
	deadlock_check.Deadlock()
}

func TypeSwitch() {
	type_switch.TypeSwitch()
	fmt.Println("type switch checked!")
}

func TypeAssertion() {
	type_assertion.TypeAssertion()
	fmt.Println("type assertion checked!")
}

func Function() {
	function_check.Function()
	fmt.Println("function checked!")
}

func Defer() {
	defer_check.Defer()
	fmt.Println("defer checked!")
}

func Recover() {
	recover_check.Recover()
}
