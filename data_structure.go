package main

import "fmt"

func data_structure() {
	/* array
		- fixed size
		- same variable type
	  - when passing array to function, it is passed by value, not by reference
	*/
	// declare array
	var nums [5]int

	// assign value
	nums[0] = 1
	fmt.Println(nums)

	// declare and assign value
	nums2 := [5]int{1, 2, 3, 4, 5}

	// passing array to function
	modifyArr(nums2)
	fmt.Println("array after modification", nums2)

	/* slice
		- dynamic size
	  - same variable type
	  - when passing slice to function, it is passed by reference, not by value
	*/

	// declare slice
	var nums3 []int
	fmt.Printf("nums3 : %v, len : %d, cap : %d", nums3, len(nums3), cap(nums3))
	fmt.Println()

	// declare and assign value
	nums4 := []int{1, 2, 3, 4, 5}
	fmt.Printf("nums4 : %v, len : %d, cap : %d", nums4, len(nums4), cap(nums4))
	fmt.Println()

	// using make function
	nums5 := make([]int, 5, 10) // if capacity is not specified, it will be equal to length
	fmt.Printf("nums5 : %v, len : %d, cap : %d", nums5, len(nums5), cap(nums5))
	fmt.Println()

	// append
	nums6 := []int{1, 2, 3, 4, 5}
	nums6 = append(nums6, 6, 7, 8, 9, 10)

	// slicing an array or slice
	fmt.Println(nums6[1:3]) // [2, 3]

	// copy
	nums7 := make([]int, 3)
	copy(nums7, nums6)
	fmt.Println(nums7)

	// passing slice to function
	modifySlice(nums7)
	fmt.Printf("slice after modification %v", nums7)
	fmt.Println()
	// maps

	// declare map
	var m map[string]int
	fmt.Println(m)

	// declare and assign value
	m2 := map[string]int{"A": 1, "B": 2, "C": 3}
	fmt.Println(m2)

	// using make function
	m3 := make(map[string]int)
	m3["A"] = 1

	// access value
	println(m3["A"])

	// delete value
	delete(m3, "A")

	// check if key exists
	if val, ok := m3["A"]; ok {
		println(val)
	} else {
		println("key not found")
	}

	// iterate over map
	for key, value := range m2 {
		fmt.Println(key, value)
	}
}

func modifyArr(arr [5]int) [5]int {
	arr[0] = 100
	fmt.Println("inside function", arr)
	return arr
}
func modifySlice(arr []int) []int {
	arr[0] = 100
	fmt.Println("inside function", arr)
	return arr
}
