package main

import "fmt"

func main() {
	nums := []int{3,4,5,1,2}
	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	start := 0
	end := len(nums) - 1

	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] < nums[end] {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	}
	return nums[end]
}
