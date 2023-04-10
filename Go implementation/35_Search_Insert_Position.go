/*
Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.
*/
package binary_search

func searchInsert(nums []int, target int) int {
	lower := 0
	upper := len(nums) - 1
	mid := 0

	for lower <= upper {
		mid = (upper + lower) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			upper = mid - 1
		} else {
			lower = mid + 1
		}
	}

	if upper > lower {
		return upper
	} else {
		return lower
	}
}
