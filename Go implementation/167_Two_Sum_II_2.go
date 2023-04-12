/*
Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution. You may not use the same element twice.

Your solution must use only constant extra space.
*/

package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	var low int
	high := len(numbers) - 1
	var mid int

	for i := 0; i < len(numbers)-1; i++ {
		low = i + 1

		for low <= high {
			mid = (low + high) / 2

			if (numbers[mid] + numbers[i]) == target {
				return []int{i + 1, mid + 1}
			}

			if (numbers[mid] + numbers[i]) > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return nil
}

func main() {

	nums := []int{2, 7, 11, 15}
	fmt.Printf("%#v\n", twoSum(nums, 9))

	nums = []int{2, 3, 4}
	fmt.Printf("%#v\n", twoSum(nums, 6))

	nums = []int{-1, 0}
	fmt.Printf("%#v\n", twoSum(nums, -1))

	nums = []int{5, 25, 75}
	fmt.Printf("%#v\n", twoSum(nums, 100))

	nums = []int{12, 13, 23, 28, 43, 44, 59}
	fmt.Printf("%#v\n", twoSum(nums, 66))

	nums = []int{12, 13, 23, 28, 43, 44, 59, 60, 61, 68, 70, 86, 88, 92, 124, 125, 136, 168, 173, 173, 180, 199, 212, 221, 227, 230, 277, 282, 306, 314, 316, 321, 325, 328, 336, 337, 363, 365, 368, 370, 370, 371, 375, 384, 387, 394, 400, 404, 414, 422, 422, 427, 430, 435, 457, 493, 506, 527, 531, 538, 541, 546, 568, 583, 585, 587, 650, 652, 677, 691, 730, 737, 740, 751, 755, 764, 778, 783, 785, 789, 794, 803, 809, 815, 847, 858, 863, 863, 874, 887, 896, 916, 920, 926, 927, 930, 933, 957, 981, 997}
	fmt.Printf("%#v\n", twoSum(nums, 66))
}
