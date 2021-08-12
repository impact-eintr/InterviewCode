package _413

func NumberOfArithmeticSlices(nums []int) int {
	return numberOfArithmeticSlices(nums)
}

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	d, count, ans := nums[0]-nums[1], 0, 0

	for i := 2; i < n; i++ {
		if nums[i-1]-nums[i] == d {
			count++
		} else {
			d, count = nums[i-1]-nums[i], 0
		}
		ans += count
	}
	return ans
}
