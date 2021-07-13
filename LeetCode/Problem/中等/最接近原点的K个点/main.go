package main

import "fmt"
import "sort"

func main() {
	points := [][]int{{1, 3}, {-2, 2}}
	K := 1
	fmt.Println(kClosest(points, K))
}

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:K]
}
