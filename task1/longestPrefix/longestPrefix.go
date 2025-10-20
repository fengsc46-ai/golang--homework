package main

import (
	"fmt"
	"sort"
)

type StrArr []string

func (strArr StrArr) Len() int {
	return len(strArr)
}
func (strArr StrArr) Less(i int, j int) bool {
	return strArr[i] < strArr[j]
}
func (strArr StrArr) Swap(i int, j int) {
	strArr[i], strArr[j] = strArr[j], strArr[i]
}

/*
*
查找字符串数组中的最长公共前缀,如果不存在公共前缀，返回空字符串 ""。
*/
func commonPrefix(strArr StrArr) string {
	if len(strArr) == 0 {
		return ""
	}

	sort.Sort(strArr)
	prefix := ""
	tempPrefix := ""
	tempStr := strArr[0]
	// 遍历数组，逐个字符比较，找到最长公共前缀
	for i := 1; i < len(strArr); i++ {
		for j := 0; j < len(tempStr); j++ {
			// 比较字符串数组中的每个字符串的第 j 个字符
			if tempStr[j] == strArr[i][j] {
				tempPrefix += string(tempStr[j])
				fmt.Println(tempPrefix)
				// 如果 tempPrefix 长度大于 prefix 长度，更新 prefix
				if len(tempPrefix) > len(prefix) {
					prefix = tempPrefix
				}
			} else {
				// 如果 tempPrefix 长度等于 prefix 长度，说明找到公共前缀，退出循环，更新 tempStr 和 tempPrefix
				tempStr = strArr[i]
				tempPrefix = ""
				break
			}
		}
	}
	return prefix
}

func main() {
	strArr := []string{"filower", "flow", "flight"}
	prefix := commonPrefix(strArr)
	println(prefix)
}
