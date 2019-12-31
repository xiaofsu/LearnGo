package main

import (
	"fmt"
	"strconv"
)

func main() {
	fplus := func(a int, b int) int { return a + b }
	fmt.Printf("The Number is :%d \n", fplus(1, 2))
	number := func(s string, num int) int {
		v, e := strconv.Atoi(s)
		if e == nil {
			return v + num
		}
		return 0
	}("12", 32)
	fmt.Printf("The Number is :%d \n", number)
}
