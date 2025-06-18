package main

func findMaxAverage(nums []int, k int) float64 {
	maxSum := 0
	sum := 0

	for i := 0; i < k; i++ {
		maxSum += nums[i]
	}
	sum = maxSum
	left := k

	for left < len(nums) {
		sum += nums[left] - nums[left-k]

		if sum > maxSum {
			maxSum = sum
		}
		left++
	}
	return float64(maxSum) / float64(k)
}
