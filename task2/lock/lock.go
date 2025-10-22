package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 线程安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 增加计数
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// 获取当前计数
func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

/**
 * @author:馮
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/

/**
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/

func main() {
	fmt.Println("================ question 1 ========")
	counter := SafeCounter{}
	// 启动100个goroutine同时增加计数
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}
	// 等待一段时间确保所有goroutine完成
	time.Sleep(time.Second)

	// 输出最终计数
	fmt.Printf("Final count: %d\n", counter.GetCount())

	fmt.Println("================ question 2 ========")
	var counters int64
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {

		go func() {
			defer wg.Done()
			atomic.AddInt64(&counters, 1)
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counters) // 注意：这里可能不是最终的计数，因为goroutine可能还没执行完
}
