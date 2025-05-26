package main

func longestSubarray(nums []int) int {
	max := 0
    ini := 0
	z := 0

	for end := range nums {
		if nums[end] == 0 {
			z += 1
		}
		if z > 1 {
			if nums[ini] == 0 {
				z--
			}
			ini ++
		}
		if max < end-ini{
            max = end-ini
        }
	}
	return max
}
//asdasd