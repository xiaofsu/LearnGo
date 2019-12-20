package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	str := "Hello World!"
	prefix := "H"
	suffix := "!"
	//判断字符串 str 是否以 prefix 开头
	fmt.Println(strings.HasPrefix(str, prefix)) //输出 true
	//判断字符串 str 是否以 suffix 结尾
	fmt.Println(strings.HasSuffix(str, suffix)) //输出 true

	contains()
	index()
	index2()
	replace()
	count()
	repeat()
	toLowUpper()
	trim()
	split()
	join()
	atoiItoa()
}

func contains() {
	str := "Hello World!"
	substr := "ello W"

	fmt.Println(strings.Contains(str, substr))
}

func index() {
	str := "Hello World!"
	substr1 := "oW"
	substr2 := "o W"
	fmt.Println(strings.Index(str, substr1))
	fmt.Println(strings.Index(str, substr2))
}

func index2() {
	str := "Hello World!"
	substr1 := "o"
	substr2 := "or"
	fmt.Println(strings.LastIndex(str, substr1)) //输出 7
	fmt.Println(strings.LastIndex(str, substr2)) //输出 7
}

func replace() {
	str := "Hello World!"
	substr1 := "o"
	substr2 := "A"
	s := strings.Replace(str, substr1, substr2, -1)
	fmt.Println(s) // 输出 HellA WArld!

	s2 := strings.Replace(str, substr1, substr2, 1)
	fmt.Println(s2) // 输出 HellA World!
}

func count() {
	var str string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H")) //Number of H's in Hello, how is it going, Hugo? is: 2

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg")) //Number of double g's in gggggggggg is: 5
}

func repeat() {
	str := "Hello World!"
	// 输出 Hello World!Hello World!Hello World!Hello World!Hello World!
	fmt.Println(strings.Repeat(str, 5))
}
func toLowUpper() {
	str := "你好，世界！Hello World!"
	fmt.Println(strings.ToLower(str)) //输出   你好，世界！hello world!
	fmt.Println(strings.ToUpper(str)) //输出   你好，世界！HELLO WORLD!
}

func trim() {
	str := "cut你好，世界！Hello World!cut"
	str2 := " 你好，世界！Hello World! "             //留意有空白字符
	fmt.Println(strings.Trim(str, "cut"))      //你好，世界！Hello World!
	fmt.Println(strings.TrimLeft(str, "cut"))  //你好，世界！Hello World!cut
	fmt.Println(strings.TrimRight(str, "cut")) //cut你好，世界！Hello World!

	fmt.Println(strings.TrimSpace(str2)) //你好，世界！Hello World!
}
func split() {
	str := "cutHello World!cut"
	splitStr := strings.Split(str, "o")
	for i, s := range splitStr {
		fmt.Printf("第 %d 位是：%s \n", i, s)
		/*
			第 0 位是：cutHell
			第 1 位是： W
			第 2 位是：rld!cut
		*/
	}

	fields := strings.Fields(str)
	for i, s := range fields {
		fmt.Printf("第 %d 位是：%s \n", i, s)
		/*
			第 0 位是：cutHello
			第 1 位是：World!cut
		*/
	}
}

func join() {

	str2 := "My Name is xiaoFsu!"
	sl2 := strings.Fields(str2)
	fmt.Printf("Splitted in slice: %v\n", sl2) //Splitted in slice: [My Name is xiaoFsu!]
	for _, val := range sl2 {
		fmt.Printf("%s - ", val) //My - Name - is - xiaoFsu! -
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3) //sl2 joined by ;: My;Name;is;xiaoFsu!
}

func atoiItoa() {
	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
