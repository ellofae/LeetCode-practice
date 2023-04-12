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

	mergeSort(nums, 0, len(nums)-1)
	return nums
}

// Bubble sort
// Algorithm's average and worst time complexity: O(n^2)
// Algorithm's best time complexity: O(n)
func mergeSort(nums []int, low int, high int) {
	if low < high {
		mid := (high + low) / 2 // low + (high-low)/2 //

		mergeSort(nums, low, mid)
		mergeSort(nums, mid+1, high)

		merge(nums, low, mid, high)
	}
}

func merge(nums []int, low, mid, high int) {
	temp_1 := make([]int, mid-low+1)
	temp_2 := make([]int, high-mid)

	for i := 0; i < len(temp_1); i++ {
		temp_1[i] = nums[low+i]
	}

	for j := 0; j < len(temp_2); j++ {
		temp_2[j] = nums[mid+j+1]
	}

	i := 0
	j := 0
	k := low
	for i < len(temp_1) && j < len(temp_2) {
		if temp_2[j] >= temp_1[i] {
			nums[k] = temp_1[i]
			i++
		} else {
			nums[k] = temp_2[j]
			j++
		}
		k++
	}

	for i < len(temp_1) {
		nums[k] = temp_1[i]
		i++
		k++
	}

	for j < len(temp_2) {
		nums[k] = temp_2[j]
		j++
		k++
	}
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	inOrder := sortedSquares(nums)

	fmt.Printf("%#v\n", inOrder)
}
