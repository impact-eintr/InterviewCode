package main

import "fmt"
import "sort"

func main() {
	num := 15
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("vim-go")
	fmt.Println(onesCount(num))
	fmt.Println(sortByBits(nums))
}

func onesCount(x int) (c int) {
	for ; x > 0; x = x >> 1 {
		c += x % 2
	}
	return
}

func sortByBits(a []int) []int {
	sort.Slice(a, func(i, j int) bool {
		x, y := a[i], a[j]
		cx, cy := onesCount(x), onesCount(y)
		return cx < cy || cx == cy && x < y

	})
	return a

}

//func sortByBits(s []int) []int{
//
//}
