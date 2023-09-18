package main

import "fmt"

func removeElement(nums []int, val int) int {
	length := len(nums)
	res := 0
	for i := 0; i < length; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	nums = nums[:res]
	return res
}

func main() {
	var nums1 = []int{3, 2, 2, 3}
	var val1 = 3
	var result1 = removeElement(nums1, val1)
	fmt.Println(result1)
	var nums2 = []int{0, 1, 2, 2, 3, 0, 4, 2}
	var val2 = 2
	var result2 = removeElement(nums2, val2)
	fmt.Println(result2)
}
