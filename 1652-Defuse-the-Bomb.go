package main

func decrypt(code []int, k int) []int {
    n := len(code)
	ret := make([]int, n)

	if k == 0 {
		return ret
	}

	left, right := 1, k+1
	
	if k < 0 {
		left, right = n + k, n
	}

	sum := 0
	for i := left; i < right; i++ {
		sum += code[i]
	}

	for i := range code {
		ret[i] = sum
		sum = sum - code[(left + i) % n] + code[(right + i) % n]
	}

	return ret
}