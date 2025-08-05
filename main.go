package main

import "fmt"

func main() {
	mainCopy()
	nums := []int{5,7,1,4}

	fmt.Println(decrypt(nums, 3))
}