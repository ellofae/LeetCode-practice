/*
Given an integer array nums sorted in non-decreasing order,
return an array of the squares of each number sorted in non-decreasing order.
*/

package main

import "fmt"

func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}

	return quickSort(nums, 0, len(nums))
}

// Quick sort (Быстрая сортировака)
// Algorithm's worst time complexity: O(n^2)
// Algorithm's average time complexity: O(n*log(n))
// Algorithm's best time complexity: O(n)

func quickSort(nums []int, l int, r int) []int {
	if l < r {
		nums, piviotIndex := partion(nums, l, r)

		nums = quickSort(nums, l, piviotIndex)
		nums = quickSort(nums, piviotIndex+1, r)
	}
	return nums
}

func partion(nums []int, l int, r int) ([]int, int) {
	piviot := nums[r-1]
	i := l - 1

	for j := l; j < r-1; j++ {
		if piviot > nums[j] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[r-1] = nums[r-1], nums[i+1]

	return nums, i + 1
}

func main() {
	nums := []int{8, 2, 4, 7, 1, 3, 9, 6, 5}
	inOrder := sortedSquares(nums)

	fmt.Printf("%#v\n", inOrder)
}
