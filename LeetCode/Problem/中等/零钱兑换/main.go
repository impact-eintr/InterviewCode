package main

import "fmt"
import "math"

const (
	MAX_INT = math.MaxInt64
)

func main() {
	//coins := []int{186, 419, 83, 408}
	coins := []int{1, 2, 5}
	amount := 19
	//amount := 11
	fmt.Println(coinChange2(coins, amount))
}

func coinChange1(coins []int, amount int) int {

	memo := make(map[int]int)
	res := dp(coins, amount, memo)
	fmt.Println(memo)
	return res
}

func dp(coins []int, n int, memo map[int]int) int {

	for k := range memo {
		if k == n {
			return memo[n]
		}
	}
	if n == 0 {
		return 0
	}
	if n < 0 {
		return -1
	}
	res := MAX_INT
	for _, coin := range coins {
		subproblem := dp(coins, n-coin, memo)
		if subproblem == -1 {
			continue
		}
		res = min(res, 1+subproblem)
	}
	if res != MAX_INT {
		memo[n] = res
	} else {
		return -1
	}
	return memo[n]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for index := range dp {
		dp[index] = amount + 1
	}
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
		fmt.Println(dp)
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
