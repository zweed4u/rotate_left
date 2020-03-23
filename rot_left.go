package main

import (
	"fmt"
)

func rotateLeft(arr []int, rotations int) []int {
	// make copy of the array
	rotatedArray := make([]int, 0)
	rotatedArray = append(rotatedArray, arr...)

	if rotations == 0 {
		return rotatedArray
	}
	lengthOfArray := len(arr)

	// reduce multiple full array loop to 1 pass
	rotations %= lengthOfArray

	for index := range arr {
		// look ahead in the array to the index whose value will 
		// be rotated to current index and wrap around len array 
		// (module) ensuring we don't go out of bounds by adding
		// the rotations to the index 
		rotatedIndex := (index + rotations) % lengthOfArray
		rotatedArray[index] = arr[rotatedIndex]
	}
	return rotatedArray
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	d := 2
	fmt.Println("Original array:", a)
	newArray := rotateLeft(a, d)
	fmt.Println("Rotated left", d, "times:", newArray)
}
