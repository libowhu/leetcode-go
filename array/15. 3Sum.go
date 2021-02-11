func threeSum(nums []int) [][]int {
	var result [][]int
	duplicate := map[[3]int]bool{}
	sort.Ints(nums)
	prevNum := math.MaxInt32
	for i, v := range nums {
		if v == prevNum {
			continue
		}
		prevNum = v
		temp := twoSum(nums[i+1:], -v)
		for _, arr := range temp {
			if _, ok := duplicate[[3]int{arr[0], arr[1], v}]; !ok {
				result = append(result, append(arr, v))
				duplicate[[3]int{arr[0], arr[1], v}] = true
			}
		}
	}
	return result
}

func twoSum(nums []int, target int) [][]int {
	var result [][]int
	lookup := map[int]bool{}
	for _, v := range nums {
		_, ok := lookup[target-v]
		if ok {
			result = append(result, []int{target - v, v})
		} else {
			lookup[v] = true
		}
	}
	return result
}

func threeSum(nums []int) [][]int {
	var result [][]int
	duplicate := map[[3]int]bool{}
	sort.Ints(nums)
	prevNum := math.MaxInt32
	for i, v := range nums {
		if v == prevNum {
			continue
		}
		prevNum = v
		temp := twoSum(nums[i+1:], -v)
		for _, arr := range temp {
			if _, ok := duplicate[[3]int{arr[0], arr[1], v}]; !ok {
				result = append(result, append(arr, v))
				duplicate[[3]int{arr[0], arr[1], v}] = true
			}
		}
	}
	return result
}

func twoSum(nums []int, target int) [][]int {
	var result [][]int
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l]+nums[r] == target {
			result = append(result, []int{nums[l], nums[r]})
			for r > 1 && nums[r] == nums[r-1] {
				r -= 1
			}
			for l < len(nums)-1 && nums[l] == nums[l+1] {
				l += 1
			}
			l += 1
		} else if nums[l]+nums[r] > target {
			r -= 1
		} else {
			l += 1
		}
	}
	return result
}


