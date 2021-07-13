package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	res, old := 0, x
	//先求出反转值
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
	}
	return res == old|res == 0

}
