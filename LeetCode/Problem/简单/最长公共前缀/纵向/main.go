package main

import "fmt"

func main() {
	test := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(test))
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""

	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]

			}

		}

	}
	return strs[0]

}
