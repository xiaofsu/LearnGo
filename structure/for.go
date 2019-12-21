package main

import "fmt"

func main() {
	//for i:=0;i<5;i++{
	//	fmt.Printf("The Number is : %d \n",i)
	//}
	test4()
}

func test3() {
	for i, j := 0, 5; i < j; i, j = i+1, j-1 {
		fmt.Printf("The Number i is: %d \n", i)
		fmt.Printf("The Number j is: %d \n", j)
	}
}

func test4() {
	var a [5]int
	for idx, val := range a {
		fmt.Printf("The idx is :%d, The Val is :%d \n", idx, val)
	}

	var intP *int
	for _, val := range a {
		intP = &val
		*intP = 23
		fmt.Println(val)
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}

	for i, j, s := 0, 5, "a"; i < 33 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}

}
