/*
Given an integer array nums sorted in non-decreasing order,
return an array of the squares of each number sorted in non-decreasing order.
*/

package main

import "fmt"

func sortedSquares(nums []int) []int {
	//for i := 0; i < len(nums); i++ {
	//	nums[i] = nums[i] * nums[i]
	//	}

	return insertion_sort(nums)
}

// Insertion sort algorithm with binary search (Сортировка вставками с бинарным поиском)
// Algorithm's average and worst time complexity: O(n*log(n))
// Algorithm's best time complexity: O(n)
func insertion_sort(nums []int) []int {
	var low int
	var high int
	var mid int

	for i := 1; i < len(nums); i++ {
		j := i - 1
		key := nums[i]
		location := 0

		if key < nums[j] {
			low = 0
			high = i - 1

			for low <= high {
				mid = (low + high) / 2

				if key >= nums[mid] && key <= nums[mid+1] {
					location = mid + 1
					break
				}

				if key >= nums[mid] {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}

			for j >= location {
				nums[j+1] = nums[j]
				j--
			}
			nums[j+1] = key
		}
	}
	return nums
}

func main() {
	//nums := []int{-4, -1, 0, 3, 10}
	//nums := []int{1, 2, 6, 3, 8, 7}
	//nums := []int{1, 2, 6, 3, 8, 7}
	//nums := []int{-4, -4, -3}
	//nums := []int{-4, -3, -3, -2, 2} // 16, 9, 9, 4, 4
	nums := []int{4, 1, 2, 6, 5, 5, 11, 9, 7, 12, 8}
	inOrder := sortedSquares(nums)

	fmt.Printf("%#v\n", inOrder)
}
