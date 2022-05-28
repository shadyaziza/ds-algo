package leetcode

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := 0
	for _, v := range nums {
		currentSum += v
		if currentSum > maxSum {
			maxSum = currentSum
		}
		if currentSum < 0 {
			currentSum = 0
		}

	}
	return maxSum
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
