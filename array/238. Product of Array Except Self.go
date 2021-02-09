func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			left[i] = 1
		} else {
			left[i] = left[i-1] * nums[i-1]
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			right[i] = 1
		} else {
			right[i] = right[i+1] * nums[i+1]
		}
	}
	for i := 0; i < n; i++ {
		result[i] = left[i] * right[i]
	}
	return result
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	output := make([]int, n)
	for cum, i := 1, 0; i < n; i++ {
		output[i] = cum
		cum *= nums[i]
	}
	for cum, i := 1, n-1; i >= 0; i-- {
		output[i] *= cum
		cum *= nums[i]
	}
	return output
}