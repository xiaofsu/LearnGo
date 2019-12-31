package main

import "fmt"

func main() {
	i, _, j := testFunction6()
	fmt.Printf("The number i is :%d,The number j is :%d \n", i, j)
}

func testFunction6() (i int, _ int, j int) {
	return 1, 3, 4
}
