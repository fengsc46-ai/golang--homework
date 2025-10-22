package main

import (
	"fmt"
	"time"
)

/*
*
编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
*/
//func sendOnly(ch chan<- int) {
//	for i := 0; i < 10; i++ {
//		ch <- i
//		fmt.Printf("生成: %d\n", i)
//	}
//	close(ch)
//}
//
//func recOnly(ch <-chan int) {
//	for i := range ch {
//		fmt.Printf("接收: %d\n", i)
//	}
//}

/*
*题目1与题目2类似
实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
*/
func producer(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Printf("生产者发送: %d\n", i)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for i := range ch {
		fmt.Printf("消费者接收: %d\n", i)
	}
}

func main() {
	ch := make(chan int, 5)

	go producer(ch)

	go consumer(ch)

	timeout := time.After(2 * time.Second)

	select {
	case v, ok := <-ch:
		if !ok {
			fmt.Println("通道已关闭")
		}
		fmt.Printf("接收到的值: %d\n", v)
	case <-timeout:
		fmt.Println("超时")
	default:
		fmt.Println("通道为空")
		time.Sleep(2 * time.Second)
	}
}
