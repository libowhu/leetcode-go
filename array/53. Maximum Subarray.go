func maxSubArray(nums []int) int {
	res, cum := nums[0], nums[0]
	for _, v := range nums[1:] {
		if cum < 0 {
			cum = v
		} else {
			cum += v
		}

		if cum > res {
			res = cum
		}
	}
	return res
}
