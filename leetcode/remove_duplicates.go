package leetcode

import "fmt"

func Remove() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	m := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		} else {
			m[nums[i]]++
		}
	}

	return len(m)
}
