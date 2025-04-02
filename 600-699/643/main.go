package main

func findMaxAverage(nums []int, k int) float64 {
	maxSum := 0
	currentSum := 0
	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}
	maxSum = currentSum

	for i := k; i < len(nums); i++ {
		currentSum += nums[i]
		currentSum -= nums[i-k]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return float64(maxSum) / float64(k)
}
