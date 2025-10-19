package main

import "fmt"

/*
*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/
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
