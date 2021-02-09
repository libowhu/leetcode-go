// https://leetcode.com/problems/two-sum/
package main

import "fmt"

// O(n2)
// func twoSum(nums []int, target int) []int {
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums); j++ {
// 			if i != j && nums[i]+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// O(n)
func twoSum(nums []int, target int) []int {
	lookup := make(map[int]int)
	for i, v := range nums {
		j, ok := lookup[target-v]
		if ok {
			return []int{j, i}
		}
		lookup[v] = i
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4, 5}, 5))
}
