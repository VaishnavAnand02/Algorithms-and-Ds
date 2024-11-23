package main

import (
	"fmt"
)

func main() {
	array := getInp()
	fmt.Println("Unsorted Array is:", array)
	array = SelectionSort(array)
	fmt.Println("Sorted Array is :", array)
}
func getInp() []int {
	var n int
	fmt.Println("Enter the number of elements")
	fmt.Scan(&n)
	a := make([]int, n)

	fmt.Println("Enter the array elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	return a
}

// Selection sort takes the first index of unsorted array and swaps it with min element index
func SelectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		temp := a[i]
		a[i] = a[min]
		a[min] = temp
	}
	return a
}
