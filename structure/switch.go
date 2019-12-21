package main

import "fmt"

func main() {
	i := 1
	switch i {
	case 0:
		fmt.Println("Hi, The Number is 0")
		fallthrough
	case 1:
		fmt.Println("Hi, The Number is 1") //Hi, The Number is 1
		fallthrough
	case 2:
		fmt.Println("Hi, The Number is 2") //Hi, The Number is 2
	case 3:
		fmt.Println("Hi, The Number is 3")
	default:
		fmt.Println("Hi,Default")
	}
	test2()
}

func test2() {
	i := 1
	switch {
	case (i + 1) == 2:
		fmt.Println("222")
	case (i - 1) == 0:
		fmt.Println("000")
	}
}
