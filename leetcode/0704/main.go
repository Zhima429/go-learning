package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + (right-left)>>1
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			return middle
		}
	}
	return -1
}

func main() {
	var nums1 = []int{-1, 0, 3, 5, 9, 12}
	var target1 = 9
	var result1 = search(nums1, target1)
	fmt.Println(result1)
	var nums2 = []int{-1, 0, 3, 5, 9, 12}
	var target2 = 2
	var result2 = search(nums2, target2)
	fmt.Println(result2)
}
