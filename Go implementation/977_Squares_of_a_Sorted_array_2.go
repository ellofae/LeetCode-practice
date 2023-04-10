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

	return shaker_sort(nums)
}

// Shaker sort algorithm (Шейкерная соритровака)(Соритровка перемешиванием)
// Algorithm's average and worst time complexity: O(n^2)
// Algorithm's best time complexity: O(n)
func shaker_sort(nums []int) []int {
	ident := true

	for i := 1; i < len(nums); i++ {
		ident = true
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				ident = false
			}
		}

		if ident {
			break
		}

		ident = true
		for j := len(nums) - 1; j >= 1; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
				ident = false
			}
		}

		if ident {
			break
		}
	}

	return nums
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	inOrder := sortedSquares(nums)

	fmt.Printf("%#v\n", inOrder)
}
