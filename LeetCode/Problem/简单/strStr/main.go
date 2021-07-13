package main

import "fmt"

func main() {
	test()
}

func test() {
	str := "mississippi"
	str2 := "issip"

	strStr2(str, str2)
}

//双指针(待解)
func strStr2(haystack string, needle string) int {
	haystacklen := len(haystack)
	needlelen := len(needle)
	if needle == "" {
		return 0

	}
	for i := 0; i < haystacklen-needlelen+1; i++ {
		if haystack[i] == needle[0] {
			if haystack[i:i+needlelen] == needle {
				fmt.Println(i)
				return i
			}
		}
	}
	return -1
}

//滚动哈希

//暴力
func strStr(haystack string, needle string) int {
	haystacklen := len(haystack)
	needlelen := len(needle)

	for i := 0; i < haystacklen-needlelen+1; i++ {
		fmt.Println(i)
		if haystack[i:i+needlelen] == needle {
			return i
		}
	}
	return -1
}
