package CD13

import (
	"errors"
	"fmt"
)

type Stack interface {
	Push(int)
	Pop() (int, error)
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

func (s *stack) Pop() (int, error) {
	if !s.Empty() {
		res := s.Top()
		s.top--
		return res, nil
	}
	return 0, errors.New("栈空")
}

func (s *stack) Top() int {
	return s.data[s.top]
}

func (s *stack) Empty() bool {
	return s.top == -1
}

func StackSort() {
	var N, tmp int
	fmt.Scanf("%d", &N)

	stack := new(N)
	helper := new(N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &tmp)
		stack.Push(tmp)
	}

	for !stack.Empty() {
		cur, err := stack.Pop()
		if err != nil {
			fmt.Println(err)
			return
		}

		for !helper.Empty() && helper.Top() < cur {
			tmp, _ = helper.Pop()
			stack.Push(tmp)
		}
		helper.Push(cur)
	}

	for !helper.Empty() {
		tmp, _ = helper.Pop()
		stack.Push(tmp)
	}

	for !stack.Empty() {
		tmp, _ = stack.Pop()
		fmt.Printf("%d ", tmp)
	}

}
