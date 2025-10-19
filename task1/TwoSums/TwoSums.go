package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i, _ := range nums {
		for j, _ := range nums {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
