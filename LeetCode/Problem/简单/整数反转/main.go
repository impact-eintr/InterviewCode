package main

import (
	"fmt"
)

func main() {
	n := 12345
	res := myreverse(n)
	fmt.Println(res)
}

func reverse(x int) int {
	y := 0
	for x != 0 {
		y = y*10 + x%10
		if !(-(1<<31) <= y && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}
