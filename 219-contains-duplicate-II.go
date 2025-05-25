package main

func containsNearbyDuplicate(nums []int, k int) bool {
    var ini int
    cache := make(map[int]int, len(nums))

    for right, curr := range nums {
        if right - ini > k {
            cache[nums[ini]]--
            ini++
        }
        cache[curr]++
        if cache[curr] > 1 {
            return true
        }
    }
    return false
}