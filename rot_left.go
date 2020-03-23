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
