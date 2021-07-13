package main

import (
	"fmt"
)

func main() {

	//A := []int{0, 2, 3, 4, 5, 2, 1, 0}
	//B := []int{0, 3, 2, 1}
	B := []int{3, 5, 5}
	fmt.Println(validMountainArray(B))
}

func validMountainArray(A []int) bool {
	length := len(A)
	if length < 3 {
		return false
	}

	i, j := 0, length-1
	for i+1 < length && A[i] < A[i+1] {
		i++
	}
	for j > 0 && A[j-1] > A[j] {
		j--
	}

	return i == j && i != 0 && j != length
}
