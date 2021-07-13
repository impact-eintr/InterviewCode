package main

import "fmt"

func main() {
	test := "[]}{({})"
	fmt.Println(isValid(test))
}

func isValid(s string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}
	stack := []byte{}
	stack = append(stack, s[0])

	for i := 1; i < length; i++ {
		stacklen := len(stack)
		if stacklen == 0 {
			stack = append(stack, s[i])
			//stack = stack[:len(stack)-1]
		} else if s[i]-stack[stacklen-1] == 1 || s[i]-stack[stacklen-1] == 2 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}

	}
	return len(stack) == 0
}
