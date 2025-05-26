package main

import "fmt"

func main() {
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))ds
	nums := []int{0, 1, 1, 1, 1, 0, 1}
	fmt.Println(longestSubarray(nums))
	fmt.Println(containsNearbyDuplicate(nums, 3))
}
