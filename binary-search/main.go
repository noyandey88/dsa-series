package binarysearch

import "math"

// O(log n) logarithmic
func BinarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := int(math.Floor(float64((start + end) / 2)))

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
