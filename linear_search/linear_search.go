package linear_search

import "fmt"

func LinearSearch() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var search = 57
	var size = len(arr)
	result := linearSearch(arr, search, size)
	if result == -1 {
		fmt.Println("Element is not present in array")
	} else {
		fmt.Println("Element is present at index", result, "and value is", arr[result])
	}
}

func linearSearch(arr []int, search int, size int) int {
	if size == 0 {
		return -1
	} else if arr[size-1] == search {
		return size - 1
	} else {
		return linearSearch(arr, search, size-1)
	}
}
