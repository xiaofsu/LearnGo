package main

import (
	"fmt"
)

func main() {

	i := 100
	fmt.Printf("An Integer:%d ,it's localtion in memory:%p \n", i, &i)
	//An Integer:100 ,it's localtion in memory:0xc00000a0e8
	//An Integer:100 ,it's localtion in memory:0xc00006e090
	//testPointer()
	//testPointer2()
	testPointer3()
}

func testPointer() {
	i := 100
	var intP *int
	intP = &i
	fmt.Printf("The 'intP' in memory :%d ,The 'i' in memory:%d \n", *intP, i) //The 'intP' in memory :824633762056 ,The 'i' in memory:824633762056
}

func testPointer2() {
	i := 100
	var intP *int
	intP = &i
	//The 'intP' is :100 ,The 'i' is:100
	fmt.Printf("The 'intP' is :%d ,The 'i' is:%d \n", *intP, i)
	i2 := *(&i)
	//The 'i2' is :100
	fmt.Printf("The 'i2' is :%d", i2)
}

func testPointer3() {
	//const i = 5
	//ptr := &i		//Cannot take the address of i
	//ptr2 := &10		//Cannot take the address of 10

	var p *int = nil
	*p = 1
	fmt.Println(*p)

	//time.Sleep
}
