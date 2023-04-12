/*
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
*/

func rotate(nums []int, k int) {
	temp_arr := make([]int, len(nums))
	length := len(nums)

	for i := range nums {
		temp_arr[(i+k)%length] = nums[i]
	}

	for i := 0; i < length; i++ {
		nums[i] = temp_arr[i]
	}
}
