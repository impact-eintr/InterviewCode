package CD1

import (
	"fmt"
)

func Find() {
	var N, M, target int
	fmt.Scanf("%d%d%d", &N, &M, &target)
	fmt.Println(N, M, target)

	tmp := 0
	flag := false
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Scanf("%d", &tmp)
			if tmp == target {
				flag = true
			}
		}
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
