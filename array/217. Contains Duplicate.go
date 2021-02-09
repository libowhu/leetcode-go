// https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	lookup := map[int]bool{}
	for _, v := range nums {
		_, ok := lookup[v]
		if ok {
			return true
		}
		lookup[v] = true
	}
	return false
}
