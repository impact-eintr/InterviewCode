package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func maxProfit(price []int) int {
	length := len(price)
	if length < 2 {
		return 0
	}

	res := 0
	for i := 1; i < length; i++ {
		diff := price[i] - price[i-1]
		if diff > 0 {
			res += diff
		}
	}
	return res
}
