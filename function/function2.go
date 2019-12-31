package main

import "fmt"

func main() {
	num := 10
	var num2, num3 int
	num2, num3 = testFunction3(num)
	fmt.Printf("The Num is :%d,The Num2 is :%d,The Num3 is :%d \n", num, num2, num3)
	num2, num3 = testFunction4(num)
	fmt.Printf("The Num is :%d,The Num2 is :%d,The Num3 is :%d \n", num, num2, num3)
	num2, num3 = testFunction5(num)
	fmt.Printf("The Num is :%d,The Num2 is :%d,The Num3 is :%d \n", num, num2, num3)
}

// 非命名返回值
func testFunction3(num int) (int, int) {
	return 2 * num, 3 * num
}

// 命名返回值
func testFunction4(num int) (x2 int, x3 int) {
	x2 = 2 * num
	x3 = 3 * num
	return // return 则会返回x2与x3
}

// 命名返回值
func testFunction5(num int) (x2 int, x3 int) {
	x2 = 2 * num
	x3 = 3 * num
	return 1, 2
}
