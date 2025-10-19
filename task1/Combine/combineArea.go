package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	start int
	end   int
}

type Intervals []Interval

func (intervals Intervals) Len() int {
	return len(intervals)
}
func (intervals Intervals) Less(i int, j int) bool {
	return intervals[i].start < intervals[j].start
}
func (intervals Intervals) Swap(i int, j int) {
	intervals[i], intervals[j] = intervals[j], intervals[i]
}

func main() {
	intervals := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}

	result := combineArea(intervals)
	fmt.Println(result)

}

func combineArea(intervals []Interval) []Interval {
	// 对区间按起始位置进行排序
	sort.Sort(Intervals(intervals))

	areas := []Interval{intervals[0]}

	for _, value := range intervals[1:] {
		lastElement := areas[len(areas)-1].end
		if value.start <= lastElement {
			// 如果有重叠，就合并它们。合并是通过更新最后一个合并区间
			// 的结束位置为这两个区间结束位置的最大值来实现的。
			areas[len(areas)-1].end = max(lastElement, value.end)
		} else {
			// 如果没有重叠，就直接把当前区间添加到合并后的数组中
			areas = append(areas, value)
		}
	}
	return areas
}
