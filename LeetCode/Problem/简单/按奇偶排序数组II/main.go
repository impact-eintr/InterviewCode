package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

//普通数组
func sortArrayByParityII1(A []int) []int {
	res := make([]int, len(A), len(A))
	j, k := 0, 1
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			res[j] = A[i]
			j += 2
		} else {
			res[k] = A[i]
			k += 2
		}
	}
	return res
}

//双指针
func sortArrayByParityII2(a []int) []int {
	for i, j := 0, 1; i < len(a); i += 2 {
		if a[i]%2 == 1 {
			for a[j]%2 == 1 {
				j += 2

			}
			a[i], a[j] = a[j], a[i]

		}

	}
	return a

}
