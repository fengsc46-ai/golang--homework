package main

import (
	"fmt"
	"strconv"
)

/*
*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/
func plusOne(digits []int) []int {
	// 先将数组转换
	digitsStr := ""
	for _, digit := range digits {
		digitsStr += fmt.Sprintf("%d", digit)
	}
	// 转换后，加一
	num, err := strconv.Atoi(digitsStr)
	if err != nil {
		fmt.Println(err)
	}
	num += 1
	// 再将结果转换为数组
	result := []int{}
	for num > 0 {
		result = append([]int{num % 10}, result...)
		num /= 10
	}
	return result
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
}
