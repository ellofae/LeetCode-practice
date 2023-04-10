/*
You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check.
Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.
You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version.
You should minimize the number of calls to the API.
*/

package binary_search

func isBadVersion(num int) bool {
	if num >= 5 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	lower := 0
	upper := n
	mid := 0
	prev := 0

	for lower <= upper {
		mid = (lower + upper) / 2

		if !isBadVersion(mid) {
			lower = mid + 1
		} else {
			upper = mid - 1
			prev = mid
		}
	}

	return prev
}
