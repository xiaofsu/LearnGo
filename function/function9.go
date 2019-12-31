package main

import "fmt"

func main() {

	func1 := test1()
	fmt.Printf("The Number is :%d \n", func1(1))
	func2 := test2(12)
	fmt.Printf("The Number is :%d \n", func2(12))
}

func test1() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func test2(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
