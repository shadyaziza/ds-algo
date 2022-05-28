package leetcode

func ContainsDuplicate(nums []int) bool {

	hashM := make(map[int]bool)

	for _, v := range nums {
		if _, ok := hashM[v]; ok {
			return true
		}
		hashM[v] = true
	}
	return false

}
