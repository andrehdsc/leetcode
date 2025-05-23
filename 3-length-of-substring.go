package main

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	max := 0
	lastIndex := make([]int, 128)

	start := 0
	for end := range n {
		current := s[end]
		if lastIndex[current] > start {
			start = lastIndex[current]
		}
		if end-start+1 > max {
			max = end - start + 1
		}
		lastIndex[current] = end + 1
	}
	return max
}

//submodule test
