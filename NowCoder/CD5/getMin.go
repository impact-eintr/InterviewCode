package CD5

import (
	"fmt"
)

type Stack interface {
	Push(int)
	Pop()
	Top()
	Empty()
}

type stack struct {
	size int
	top  int
	data []int
}

func new(size int) *stack {
	return &stack{size, -1, make([]int, size)}
}

func (s *stack) Push(elm int) {
	s.top++
	s.data[s.top] = elm
}

func (s *stack) Pop() {
	if !s.Empty() {
		s.top--
	}
}

func (s *stack) Top() int {
	return s.data[s.top]
}

func (s *stack) Empty() bool {
	return s.top == -1
}

func GetMin() {
	var N int

	fmt.Scanf("%d", &N)
	stack := new(N)
	minstack := new(N)

	var op string
	var num int

	for i := 0; i < N; i++ {
		fmt.Scanf("%s", &op)

		if op == "push" {
			fmt.Scanf("%d", &num)
			stack.Push(num)
			if minstack.Empty() || num <= minstack.Top() {
				minstack.Push(num)
			} else {
				minstack.Push(minstack.Top())
			}
		} else if op == "pop" {
			stack.Pop()
			minstack.Pop()
		} else if op == "getMin" {
			fmt.Println(minstack.Top())
		}
	}
}
