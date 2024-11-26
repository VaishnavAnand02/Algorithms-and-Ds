package main

import (
	"fmt"
)

func InputArray() []int {
	var n int
	fmt.Println("Enter the number of elements in the array")
	fmt.Scan(&n)

	array := make([]int, n)

	fmt.Println("Enter the array elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	return array
}

// start from the array at pos 1 and sort the array backward by inserting at the lowest element index
func InsertionSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1

		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j = j - 1
		}

		a[j+1] = key
	}
	return a
}
func main() {
	array := InputArray()
	fmt.Println("Unsorted array is:", array)

	array = InsertionSort(array)
	fmt.Println("Sorted array is:", array)
}
