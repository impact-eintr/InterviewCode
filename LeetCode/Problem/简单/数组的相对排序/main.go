package main

import "fmt"
import "sort"

func main() {
	fmt.Println("vim-go")
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	fmt.Println(relativeSortArray(arr1, arr2))
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	nummap := make(map[int]int, 1)
	for i, v := range arr2 {
		nummap[v] = i
	}
	fmt.Println(nummap)
	sort.Slice(arr1,
		func(i, j int) bool {
			return arr1[i] < arr1[j]
		})
	return arr1
}
