package main

import (
	"fmt"
)

func main() {
	fmt.Print("测试Map类型使用。")
	mapTest1()
	fmt.Println("")
	mapTest2()
	fmt.Println("")
	mapTest3()
	fmt.Println("")
	mapTest4()
	fmt.Println()
	mapTest5()

}

/**
* author:xiaoFsu
* time:2019.12.17 17:17:08
* leetcode题目：https://leetcode-cn.com/problems/two-sum/
* 使用java 与 go 相差10倍内存消耗233
* 几秒前		通过	4 ms	3.8 MB		Golang
* 6 个月前	通过	6 ms	38.4 MB		Java
 */
func twoSum(nums []int, target int) []int {
	intMap := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		_, isExist := intMap[diff]
		if isExist {
			return []int{intMap[diff], i}
		}
		intMap[num] = i
	}
	return nil
}

/**
* author:xiaoFsu
* time:2019.12.17 17:04:07
* map 删除键值对
 */
func mapTest5() {
	mapTest := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println("删除前...")
	for k, v := range mapTest {
		fmt.Println(k + "->" + v)
	}

	delete(mapTest, "key1")

	fmt.Println("删除后...")
	for k, v := range mapTest {
		fmt.Println(k + "->" + v)
	}

}

/**
* author:xiaoFsu
* time:2019.12.17 17:01:55
* map 进行遍历
 */
func mapTest4() {
	mapTest := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	//map 遍历时与添加键值对的顺序无关
	for k, v := range mapTest {
		fmt.Println(k + "->" + v)
	}
}

/**
* author:xiaoFsu
* time:2019.12.17 17:00:32
* map 的 判断键是否存在
 */
func mapTest3() {
	//在初始化map时如果指定参数的话，则不需要make函数
	//注意的是，需要在最后一个逗号也不能丢的哦
	mapTest := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	//检查是否存在key值
	//value 变量是，如果map中存在则取出来value值
	//isExist 如果存在则为true 否则为false
	value, isExist := mapTest["key1"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("键不存在！")
	}
}

/**
* author:xiaoFsu
* time:2019.12.17 17:00:55
* map 在初始化的时候进行赋值
 */
func mapTest2() {
	//在初始化map时如果指定参数的话，则不需要make函数
	//注意的是，需要在最后一个逗号也不能丢的哦
	mapTest := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(mapTest)
}

/**
* author:xiaoFsu
* time:2019.12.17 17:01:17
* map 进行初始化，并赋值
 */
func mapTest1() {
	//map 在创建时需要使用make 来进行创建，[string]代表的为key为string类型。int代表的value为int类型
	//8则是初始化 8个容量，非必须制定，但是最好在初始化map时指定一个容量
	mapTest := make(map[string]int, 8)
	mapTest["01"] = 01
	mapTest["02"] = 02
	fmt.Println(mapTest)
}
