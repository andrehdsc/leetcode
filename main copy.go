package main

import "fmt"

func mainCopy() []int {
	nums := []int{5,7,1,4}
	
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(longestSubarray(nums))
	fmt.Println(containsNearbyDuplicate(nums, 3))
	fmt.Println(findMaxAverage(nums, 3))
	fmt.Println(findRepeatedDnaSequences("A,B,C"))
	fmt.Println(minSubArrayLen(1, nums))
	fmt.Println(decrypt(nums, 3))

	return nums
}