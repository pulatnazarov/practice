package binary_search

import "fmt"

func BinarySearch() {
	var arr = []int{1, 2, 3, 4, 5}
	target := 60
	high := len(arr) - 1
	low := 0

	result := binarySearchRecursiveMethod(arr, target, low, high)

	if result == -1 {
		fmt.Println("Element not found")
	} else {
		fmt.Println("Element found at index: ", result, "and the value: ", arr[result])
	}

	//result := binarySearchIterationMethod(arr, target, low, high)
	//if result == -1 {
	//	fmt.Println("Element not found")
	//} else {
	//	fmt.Println("Element found at index: ", result, "and the value: ", arr[result])
	//}
}

func binarySearchIterationMethod(arr []int, target, low, high int) int {
	for low <= high {
		middle := low + (high-low)/2
		if arr[middle] == target {
			return middle
		} else if arr[middle] < target {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	return -1
}

func binarySearchRecursiveMethod(arr []int, target, low, high int) int {
	if high >= low {
		middle := low + (high-low)/2
		if arr[middle] == target {
			return middle
		} else if arr[middle] > target {
			high = middle - 1
			return binarySearchRecursiveMethod(arr, target, low, high)
		}
		return binarySearchRecursiveMethod(arr, target, low, high)
	}

	return -1
}
