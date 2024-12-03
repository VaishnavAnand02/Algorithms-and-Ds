package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter the no of elements")
	fmt.Scan(&n)

	array := make([]int, n)

	fmt.Println("Enter the array elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}
	fmt.Println("Unsorted Array:", array)
	array = MergeSort(array)
	fmt.Println("Sorted Array:", array)
}

func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) / 2
	left := a[:mid]
	right := a[mid:]

	sleft := MergeSort(left)
	sRight := MergeSort(right)

	return merge(sleft, sRight)
}
func merge(left []int, right []int) []int {
	var i, j int = 0, 0
	var result []int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...) //... is like a spread operator in javascript it converts to slice into individual elements
	result = append(result, right[j:]...)

	return result
}
