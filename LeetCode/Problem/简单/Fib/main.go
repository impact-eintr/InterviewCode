package main

import "fmt"

//动态规划 FibArray
//初始化 base case
//dp[0][0][...] = base
//进行状态转移
//for status1 := range all the status1:
//	for status2 := range all the status2:
//		for ...:
//			dp[status1][status2][status...] = MAX||MIN(select1,select2...)
//
func main() {
	fib2(10)
}

//带备忘录的递归解法
func fib1(N int) int {
	if N < 1 {
		return 0
	}
	resarray := make([]int, N+1)
	res := helper(resarray, N)
	fmt.Println(resarray[1:])
	return res
}

func helper(resarray []int, n int) int {
	if n == 1 || n == 2 {
		resarray[n] = 1
		return 1
	}
	if resarray[n] != 0 {
		return resarray[n]
	}
	resarray[n] = helper(resarray, n-1) + helper(resarray, n-2)
	return resarray[n]
}

//DP数组的迭代算法
func fib2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if n == 1 || n == 2 {
		return 1
	}
	prev, curr := 1, 1
	fmt.Print("1 1 ")
	for i := 3; i < n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
		fmt.Printf("%d ", sum)
	}
	return curr
}
