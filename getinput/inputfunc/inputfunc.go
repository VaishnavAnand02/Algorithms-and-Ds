package inputfunc

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
