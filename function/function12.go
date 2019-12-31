package main

import (
	"fmt"
	"time"
)

// 计算40个菲波那切数列
const count = 50

var fbnc [count]int

func main() {
	result := 0
	start := time.Now()
	for i := 0; i < count; i++ {
		result = fbnc1(i)
		fmt.Printf("The Fbnc %d is :%d \n", i, result)
	}
	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("Run Time is :%d \n", diff)
}
func fbnc1(f int) (ret int) {
	if fbnc[f] != 0 {
		ret = fbnc[f]
		return
	}
	if f <= 1 {
		ret = 1
	} else {
		ret = fbnc1(f-1) + fbnc1(f-2)
	}
	fbnc[f] = ret
	return
}
