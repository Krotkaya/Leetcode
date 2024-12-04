package main

import "fmt"

func runningSum(nums []int) []int {
	newArray := nums
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			newArray[i] = nums[i]
		} else {
			newArray[i] = nums[i-1] + nums[i]
		}
		fmt.Println(newArray)
	}
	return (newArray)
}

func main() {
	var array []int = []int{1, 2, 3, 44}
	runningSum(array)
}
