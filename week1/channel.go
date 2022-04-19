package main

import (
	"fmt"
	"time"
)

func ProducerAndConsumerTest() {
	// 声明通道
	ch := make(chan int, 10)

	defer close(ch)

	// 构建消费者逻辑
	// 消费者只获取
	go func(ch <-chan int) {
		timer := time.NewTicker(time.Second)
		for {
			select {
			case <-timer.C:
				num := <-ch
				fmt.Println("receive: ", num)
			}
		}
	}(ch)

	// 构建生产者逻辑
	// 生产者只存放
	func(ch chan<- int) {
		timer := time.NewTicker(time.Second)
		for {
			select {
			case <-timer.C:
				ch <- 1
				fmt.Println("produce: ", 1)
			default:
				//fmt.Println("p waiting...")
			}
		}
	}(ch)
}
