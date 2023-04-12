/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
*/

package main

import "fmt"

func rotate(nums []int, k int) {
	temp_arr := make([]int, len(nums))
	length := len(nums)

	for i := range nums {
		temp_arr[(i+k)%length] = nums[i]
	}
}

func main() {
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	//rotate(nums, 3)

	nums := []int{1, 2, 3, 4}
	rotate(nums, 2)
	fmt.Println(nums)
	//nums = []int{-1, -100, 3, 99}
	//rotate(nums, 2)
}
