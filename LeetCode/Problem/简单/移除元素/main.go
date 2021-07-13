package main

import "fmt"

func main() {
	test1 := []int{1, 1, 2, 3, 4, 2, 5, 5}
	test2 := []int{2, 2, 2, 2, 2, 2, 4}
	removeElement(test1, 2)
	removeElement(test2, 2)
}

//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {
	j, t := 0, len(nums)-1
	for ; j < len(nums); j++ {
		for t >= j && nums[t] == val {
			t--
		}
		if t >= j && nums[j] == val {
			nums[j] = nums[t]
			t--
		}
	}
	fmt.Println(nums, "  ", t+1)
	return t
}
