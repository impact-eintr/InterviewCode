package _217

//给定一个整数数组，判断是否存在重复元素。
//如果存在一值在数组中出现至少两次，函数返回 true 。如果数组中每个元素都不相同，则返回 false 。
func containsDuplicate(nums []int) bool {
	m := make(map[int]int, len(nums))
	for _, n := range nums {
		go func() {
			m[n]++
		}()
		c := m[n]
		if c > 1 {
			return true
		}
	}
	return false
}
