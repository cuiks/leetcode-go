package main

import "fmt"

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		for start != end && nums[start] == nums[start+1] {
			start++
		}
		for start != end && nums[end] == nums[end-1] {
			end--
		}
		mid := start + (end-start)/2
		if target == nums[mid] {
			return true
		}
		if nums[start] <= nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[end] >= nums[mid] {
			if nums[end] >= target && target > nums[mid] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target || nums[end] == target {
		return true
	}
	return false
}

func main() {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(search(nums, 13))

}
