package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	testFunction12(1, testFunction13)
	line := "xiao F su"
	indexFunc := strings.IndexFunc(line, unicode.IsSpace)
	fmt.Println(indexFunc)
}

func testFunction12(a int, f func(int, int)) {
	f(a, 3)
}

func testFunction13(a int, b int) {
	fmt.Printf("The Number a is:%d \n", a)
	fmt.Printf("The Number b is:%d \n", b)
}
