package main

import "fmt"

func main() {
	test := []int{1, 1, 2, 3, 4, 4, 5}
	fmt.Println(removeDuplicates(test))
}

//快慢双指针
func removeDuplicates(nums []int) int {
	i, j := 0, 1
	for ; j < len(nums); j++ {
		if j-i > 0 && nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
