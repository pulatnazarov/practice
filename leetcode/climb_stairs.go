package leetcode

import "fmt"

func ClimbStairs() {
	fmt.Println(climbStairs(45))
}

func climbStairs(n int) int {
	v := make([]int, n+1)
	v[0] = 1
	v[1] = 1
	for i := 2; i <= n; i++ {
		v[i] = v[i-1] + v[i-2]
	}

	return v[n]
}
