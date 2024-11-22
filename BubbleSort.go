package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter the number of elements in the array")
	fmt.Scan(&n)
	array := make([]int, n)

	fmt.Println("Enter the array elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println("Unsorted array is ", array)
	array = BubbleSorts(array, n)
	fmt.Println("Sorted array is ", array)
}

// bubble sort always start from the front and compares the two adjacent elements. On each iteration it reduces its end length by i.
func BubbleSorts(a []int, n int) []int {
	for i := 0; i < n; i++ {
		for j := 1; j < n-i; j++ {
			if a[j] <= a[j-1] {
				temp := a[j]
				a[j] = a[j-1]
				a[j-1] = temp
			}
		}
	}
	return a
}
