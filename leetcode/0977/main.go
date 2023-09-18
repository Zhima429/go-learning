package main

import "fmt"

func sortedSquares(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1
	ans := make([]int, n)
	for i <= j {
		lm, rm := nums[i]*nums[i], nums[j]*nums[j]
		if rm > lm {
			ans[k] = rm
			j--
		} else {
			ans[k] = lm
			i++
		}
		k--
	}
	return ans
}

func main() {
	var nums1 = []int{-4, -1, 0, 3, 10}
	var ans1 = sortedSquares(nums1)
	fmt.Println(ans1)
	var nums2 = []int{-7, -3, 2, 3, 11}
	var ans2 = sortedSquares(nums2)
	fmt.Println(ans2)
}
