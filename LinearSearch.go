package main

import (
	"fmt"
)

func main() {
	var key, n int
	fmt.Println("Enter the no of elements in the array")
	fmt.Scan(&n) //scan takes input without any format unline Scanf

	array := make([]int, n) //to make a slice of an array int of size n
	//make has parameter ([]datatype,length,capacity)
	fmt.Println("Enter the array elements")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	fmt.Println("Enter the key to be searched")
	fmt.Scan(&key)

	flag := false
	for i := 0; i < n; i++ {
		if array[i] == key {
			fmt.Printf("Key %d found at position %d\n", key, i+1)
			flag = true
			break
		}
	}
	if !flag {
		fmt.Printf("Key %d not found\n", key)
	}
}
