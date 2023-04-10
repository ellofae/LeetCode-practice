/*
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.
*/
package binary_search

func search(nums []int, target int) int {
	lower := 0
	upper := len(nums) - 1
	mid := 0

	for lower <= upper {
		mid = (upper + lower) / 2

		if target == nums[mid] {
			return mid
		}

		if target > nums[mid] {
			lower = mid + 1
		} else {
			upper = mid - 1
		}
	}

	return -1
}
