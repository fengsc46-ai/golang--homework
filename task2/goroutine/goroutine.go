package main

import (
	"fmt"
	"sync"
	"time"
)

/**
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
*/

func outPutOddNumbers() {
	for i := 1; i <= 10; i += 2 {
		fmt.Printf("%d\n", i)
		time.Sleep(10 * time.Millisecond)
	}
}
func outPutEvenNumbers() {
	for i := 2; i <= 10; i += 2 {
		fmt.Printf("%d\n", i)
		time.Sleep(10 * time.Millisecond)
	}
}

/*
*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/
func taskScheduler(tasks []func()) {
	wg := sync.WaitGroup{}
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func() {
			defer wg.Done()
			start := time.Now()
			task()
			end := time.Now()
			fmt.Printf("Task execution time: %s\n", end.Sub(start))
		}()
	}
	wg.Wait()
}

func main() {
	tasks := []func(){
		outPutOddNumbers,
		outPutEvenNumbers,
	}

	taskScheduler(tasks)
}
