func maxProduct(nums []int) int {
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		temp := 1
		for j := i; j < len(nums); j++ {
			temp *= nums[j]
			if temp > res {
				res = temp
			}
		}
	}
	return res
}

func maxProduct(nums []int) int {
	res, min, max := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		cur_min := minOf(num, num*min, num*max)
		cur_max := maxOf(num, num*min, num*max)

		if cur_max > res {
			res = cur_max
		}
		min = cur_min
		max = cur_max
	}
	return res
}

func minOf(vars ...int) int {
	min := vars[0]
	for _, v := range vars {
		if v < min {
			min = v
		}
	}
	return min
}

func maxOf(vars ...int) int {
	max := vars[0]
	for _, v := range vars {
		if v > max {
			max = v
		}
	}
	return max
}