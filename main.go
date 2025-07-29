package main

import (
	"fmt"

	twopointers "github.com/noyandey88/dsa-series/two-pointers"
)

func main() {
	// two sum with two pointers technique
	numsArr := []int{3, 5, 6, 8, 10, 14, 16}
	target := 168

	// two sum result
	left, right, err := twopointers.TwoSum(numsArr, target)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("pair is", left, right)
}
