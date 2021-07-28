package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func() {
		c <- "hello"
	}()
	msg := <-c
	fmt.Println(msg)

	//单向通道
	ch := make(chan string)

	go sc(ch)
	fmt.Println(<-ch)

	//select chan
	go speed1(c)
	go speed2(ch)
	fmt.Println("The first to arrive is :")
	select { //使用 select case 语句打印出结果，消息会通过通道发送过来，会打印出先发送过来的消息。
	case s1 := <-c:
		fmt.Println(s1)
	case s2 := <-ch:
		fmt.Println(s2)
	}

	//缓冲通道：使用缓冲通道，在缓冲区满之前接受方不会收到任何消息。
	ch1 := make(chan string, 2)
	ch1 <- "hello"
	ch1 <- "world"
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}

func sc(ch chan<- string) { //sc 是一个 Go routine，它只能向通道发送消息但不能接收消息。
	ch <- "hello chan"
}

func speed1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "speed1"
}

func speed2(ch chan string) {
	time.Sleep(time.Second)
	ch <- "speed2"
}
