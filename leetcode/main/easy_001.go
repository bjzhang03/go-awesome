package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (target == nums[i]+nums[j]) {
				result[0] = i
				result[1] = j
				return result
			}

		}
	}
	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
