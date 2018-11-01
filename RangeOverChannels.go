package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue { //使用range来访问channel的值
		fmt.Println(elem)
	}
}
