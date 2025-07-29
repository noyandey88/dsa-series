package twopointers

import "errors"

func TwoSum(arr []int, target int) (int, int, error) {
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			return arr[left], arr[right], nil
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return 0, 0, errors.New("target not found")
}
