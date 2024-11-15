package main

import (
	"fmt"
)

func main() {
	var n int
	var key int
	fmt.Println("Enter the no of elements in the array")
	fmt.Scan(&n)

	array := make([]int, n)

	fmt.Println("Enter the array elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	fmt.Println("Enter the key")
	fmt.Scan(&key)

	pos := binarySearch(array, 0, n, key)

	if pos == -1 {
		fmt.Println("Key not found")
	} else {
		fmt.Printf("Key %d found at position %d", key, pos+1)
	}
}

func binarySearch(a []int, first int, last int, key int) int {

	mid := int(first + (last-first)/2)
	if last >= first {
		if a[mid] == key {
			return mid
		}

		if key > a[mid] {
			return binarySearch(a, mid+1, last, key)
		} else {
			return binarySearch(a, first, mid-1, key)
		}
	}

	return -1
}
