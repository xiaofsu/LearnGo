package main

import "fmt"

func main() {
	bool := true
	if bool {
		fmt.Println("条件成立执行这段代码") //条件成立执行这段代码
	} else {
		fmt.Println("条件不成立执行这段代码") //没有进行输出！
	}
	fmt.Println("不管条件是否成立都会执行这段代码") //不管条件是否成立都会执行这段代码
	test1()
}

func test1() {
	max := 100
	val := 10000
	if val := 10; max > val {
		fmt.Printf("The Number Max is Big,it's :%d \n", max) //The Number Max is Big,it's :100
	} else {
		fmt.Printf("The Number val Big,it's :%d \n", val)
	}
	fmt.Println(val) //10000
}
