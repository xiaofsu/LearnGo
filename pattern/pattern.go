package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	orig := "123"
	var newStr string

	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d \n", an)
	an = an + 5
	newStr = strconv.Itoa(an)
	fmt.Printf("The new string is :%s \n", newStr)

	val, b := testPattern(12.0)
	if !b {
		fmt.Println("失败！")
		return
	}
	fmt.Println(val)

}

func testPattern(f float64) (v float64, b bool) {
	if f > 100 {
		return
	}
	return math.Sqrt(f), true
}
