package main

import "fmt"
import "time"

func main() {
	fmt.Println("vim-go")
	nums := []int{1, 2, 3, 5, 7, 9, 12, 14, 19, 23, 45}
	target := 23
	res := searchInsert(nums, target)
	fmt.Println(res)
}

func searchInsert(nums []int, target int) int {
	length := len(nums)
	l, r := 0, length-1
	resindex := length
	for l < r {
		//mid := (r-l)>>1 + l
		mid := (r + l) >> 1
		if nums[mid] < target {
			l = mid
		} else if nums[mid] > target {
			r = mid
		} else {
			resindex = mid
			break
		}
		fmt.Println(mid)
		fmt.Println(resindex)
		time.Sleep(1 * time.Second)
	}
	return resindex
}

func test() {
	a, b := 6, 7
	c := (a + b) >> 1
	fmt.Println(c)
}
