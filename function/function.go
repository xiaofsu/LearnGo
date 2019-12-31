package main

import "fmt"

func main() {
	args := "xiaoFsu"
	fmt.Printf("The old string is :%s \n", args)
	//testFunction(&args)
	//fmt.Printf("The new string is :%s \n",args)
	testFunction2(args)
	fmt.Printf("The new string is :%s \n", args)
}

func testFunction(args *string) {
	*args = "xiaofsu.com"
}
func testFunction2(args string) {
	args = "look"
}
