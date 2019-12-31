package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d \n", file, line)
	}
	where()
	fmt.Println("do  something")
	fmt.Println("do  something")
	where()
	where()
	fmt.Println("do  something")
	log.SetFlags(log.Llongfile)
	log.Print("111111111")
	test := log.Print
	test()
	end := time.Now()
	diff := end.Sub(start)
	log.Printf("The Time is :%d", diff)
	time.Sleep(1e9)
	fmt.Print("Over!")
}
