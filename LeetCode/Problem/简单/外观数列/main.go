package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	for i := 2; i <= 30; i++ {
		countAndSay(i)

	}
}

func countAndSay(n int) string {
	res := "1"
	if n == 1 {
		return res
	}
	numlist := []byte{1}
	changeloop := 0
	byteres, _, _ := change(numlist, changeloop, n-2)
	for i := 0; i < len(byteres); i++ {
		byteres[i] = byteres[i] + 48
	}
	return string(byteres)
}

func change(nums []byte, loop, Loop int) ([]byte, int, int) {
	//res := "1"
	n := len(nums)
	tempnum := nums[0]

	tmp := []byte{}
	count := byte(0)
	for i := 0; i < n; i++ {
		if nums[i] == tempnum {
			count++
		}
		for nums[i] != tempnum {
			tmp = append(tmp, count, tempnum)
			tempnum = nums[i]
			count = 1
		}
		if i == n-1 {
			tmp = append(tmp, count, tempnum)
		}
	}
	if loop < Loop {
		loop++
		tmp, _, _ = change(tmp, loop, Loop)

	}

	return tmp, loop, Loop
}
