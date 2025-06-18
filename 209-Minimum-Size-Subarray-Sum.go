package main

func minSubArrayLen(target int, nums []int) int {
	sum, ini := 0, 0
    xlen := len(nums)+1
	
	for end := range nums {
		sum = sum + nums[end]
        for sum >= target {
            if xlen > end - ini + 1 {
                xlen = end - ini + 1
            } 
            sum -= nums[ini]
            ini++
        }
	}
    if xlen == len(nums)+1{
        return 0
    }
	return xlen
}