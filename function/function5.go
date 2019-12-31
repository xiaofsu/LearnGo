package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello, There is one")
	//defer testFunction9()
	//fmt.Println("Hello, There is three")
	//i := 100
	//defer fmt.Printf("The Number is :%d \n",i)
	//for i:=0;i<5;i++{
	//	defer fmt.Printf("The Number in for is :%d \n",i)
	//}
	testFunction11()
}

func testFunction9() {
	fmt.Println("Hello, There is two")
}

func testFunction10(s string) {
	func() {
		fmt.Printf("There is :%s \n", s)
	}()
	testFunction9()
}

func testFunction11() {
	str := "xiaoFsu"
	testFunction10(str)
}
