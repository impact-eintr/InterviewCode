/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

[1,2,3]
[1,3,2]
[2,1,3]
[2,3,1]
[3,1,2]
[3,2,1]

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
*/
package main

import "fmt"

func main() {
	test := []int{2, 6, 5, 4}
	nextPermutation2(test)
	fmt.Println(test)
}

func nextPermutation1(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--

	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--

		}
		nums[i], nums[j] = nums[j], nums[i]

	}
	reverse(nums[i+1:])

}
func nextPermutation2(nums []int) {
	n := len(nums)
	i := n - 2
	j := n - 1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		k := n - 1
		for k >= 0 && nums[k] <= nums[i] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	reverse(nums[i+1:])
}
func reverse(a []int) {
	fmt.Println(a)
	for i, n := 0, len(a); i < n/2; i++ {
		fmt.Println(a[i], a[n-1-i])
		a[i], a[n-1-i] = a[n-1-i], a[i]
		fmt.Println(a[i], a[n-1-i])

	}
	fmt.Println(a)

}
