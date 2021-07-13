package main

import "fmt"

func main() {
	Num := "IX"
	num := romaToInt(Num)
	fmt.Println(num)
}

func romaToInt(s string) (res int) {
	wordmap := map[byte](int){
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	pre := 0
	for i := len(s) - 1; i >= 0; i-- {
		n := wordmap[s[i]]
		if n < pre {
			res -= n

		} else {
			res += n

		}
		pre = n

	}

	return res
}
