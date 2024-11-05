package main

import (
	"fmt"
	"math"
)

func twoLargest(nums []int) []int {
	maxValue := math.MinInt
	secondMaxValue := math.MinInt

	for _, num := range nums {
		if num > maxValue {
			secondMaxValue = maxValue
			maxValue = num
		} else if num > secondMaxValue {
			secondMaxValue = num
		}
	}
	result := []int{maxValue, secondMaxValue}
	fmt.Println(result)
	return result
}

func main() {
	twoLargest([]int{7, 8, 3})
}