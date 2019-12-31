package main

import "fmt"

func main() {
	testFunction7(1, "xiaoFsu", "OK", "No")
	slice := []int{1, 3, 4, 5, 6, 7}
	minNum := testFunction8(slice...)
	fmt.Printf("The MinNumber is :%d \n", minNum)
}

func testFunction7(num int, strN ...string) {
	fmt.Printf("The Number is :%d, The strN is:%s", num, strN)
}

func testFunction8(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	minNum := s[0]
	for _, v := range s {
		if v < minNum {
			minNum = v
		}
	}
	return minNum
}
