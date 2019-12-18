package main //声明Main包

import "fmt" //导入 fmt 包，打印字符串需要

func main() { //声明 Main 主函数
	fmt.Println("这里是 main函数。") //打印 Hello World!
}
func init() {
	fmt.Println("init 函数先于 main函数执行。")
}
