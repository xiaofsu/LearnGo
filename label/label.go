package main

import "fmt"

func main() {
LABEL1:
	for i := 0; i < 5; i++ {
		for j := 5; j >= 0; j-- {
			if j == 2 {
				fmt.Printf("The number i is:%d ,The number j is:%d \n", i, j)
				continue LABEL1
			}
		}
		fmt.Printf("The number i is :%d \n", i)
	}
	testLable()
	testLable2()
	testLable3()
}

func testLable() {
	i := 0
LABEL1:
	i++
	if i == 8 {
		fmt.Printf("The i is :%d \n", i)
		return
	}
	goto LABEL1
}

func testLable2() {
	j := 0
LABEL1:
	j++
	i := 2
	fmt.Printf("The i is :%d,The j is:%d \n", i, j)
	if j == 5 {
		return
	}
	goto LABEL1
}

func testLable3() {
	//	a := 1
	//	goto TARGET // compile error
	//	b := 9
	//TARGET:
	//	b += a
	//	fmt.Printf("a is %v *** b is %v", a, b)
}
