package main

import "fmt"

func main() {
	func3 := test3()
	fmt.Printf("The Number is :%d \n", func3(1))
	fmt.Printf("The Number is :%d \n", func3(10))
	fmt.Printf("The Number is :%d \n", func3(100))

	fbnc := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Printf("The fibonacci is :%d \n", fbnc())
	}

}

func test3() func(int) int {
	var x int
	fmt.Printf("The x is :%d \n", x)
	return func(delta int) int {
		x += delta
		return x
	}
}

func fibonacci() func() int {
	a := -1
	b := 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
