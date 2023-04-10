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

	return insertion_sort(nums)
}

// Insertion sort algorithm (Сортировка вставками)
// Algorithm's average and worst time complexity: O(n^2)
// Algorithm's best time complexity: O(n)
func insertion_sort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1

		for j >= 0 && key < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}

	return nums
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	inOrder := sortedSquares(nums)

	fmt.Printf("%#v\n", inOrder)
}
