package main

import (
	"fmt"
	"runtime"
)

type (
	testOK string
)

func Random() int {
	return 2019 * 32
}

func main() {
	testFile := test()
	fmt.Println(testFile)
	fmt.Println(myName)
	fmt.Printf("%s", runtime.Version())
}

func init() {
	fmt.Println("加载inti!")
}
