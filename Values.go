package main

import "fmt"

/**
go支持的数据类型包括String，Integer，float，Boolean...
*/
func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0 / 3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
