package main

func mainCopy() []int {
	nums := []int{5, 7, 1, 4}

	lengthOfLongestSubstring("abcabcbb")
	longestSubarray(nums)
	containsNearbyDuplicate(nums, 3)
	findMaxAverage(nums, 3)
	findRepeatedDnaSequences("A,B,C")
	minSubArrayLen(1, nums)
	decrypt(nums, 3)

	return nums
}
