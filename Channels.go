package main

import "fmt"

/**
Channels 是连接并发goroutines的管道，可以通过Channels从一个goroutine传输数据到另外一个goroutine
*/
func main() {
	//声明channels并指定传输类型
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages //<-channel：从channel接收值
	fmt.Println(msg)
}
