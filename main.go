package main

import (
	"fmt"

	slidingwindow "github.com/noyandey88/dsa-series/sliding-window"
	twopointers "github.com/noyandey88/dsa-series/two-pointers"
)

func main() {
	// two sum with two pointers technique
	numsArr := []int{3, 5, 6, 8, 10, 14, 16}
	target := 16

	// two sum result
	left, right, err := twopointers.TwoSum(numsArr, target)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("pair is", left, right)
	}

	// sliding window
	watchHistory := []int{1, 2, 3, 6, 11, 12, 14, 18, 20, 22, 25}
	maxTime := 4

	maxWatchedTime := slidingwindow.MaxWatchTime(watchHistory, maxTime)

	fmt.Println(maxWatchedTime)
}
