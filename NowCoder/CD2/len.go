package CD2

import (
	"fmt"
	"sort"
)

func Len() {
	var N, tmp int

	fmt.Scanf("%d", &N)
	lenMap := make(map[int]struct{}, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &tmp)
		lenMap[tmp] = struct{}{}
	}

	lenArr := make([]int, len(lenMap))
	index := 0
	for k, _ := range lenMap {
		lenArr[index] = k
		index++
	}

	sort.Slice(lenArr, func(i, j int) bool {
		return lenArr[i] < lenArr[j]
	})

	temp, count := 1, 1

	for i := 1; i < len(lenArr); i++ {
		if lenArr[i]-lenArr[i-1] == 1 {
			temp++
		} else {
			temp = 1
		}

		count = func() int {
			if count > temp {
				return count
			} else {
				return temp
			}
		}()
	}
	if lenArr[len(lenArr)-1]-lenArr[0]+1 != len(lenArr) {
		fmt.Println(1)
	} else {
		fmt.Println(count)
	}

}
