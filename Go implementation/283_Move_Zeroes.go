package main

// Algorithm's complexity: O(n)
func moveZeroes(nums []int) {
	count := 0
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
			continue
		}

		nums[j] = nums[i]
		j++
	}

	for i := 0; i < count; i++ {
		nums[j] = 0
		j++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 0, 12, 0, 0, 5}
	moveZeroes(nums)
}
