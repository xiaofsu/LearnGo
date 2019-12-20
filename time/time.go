package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)                                                   //输出当前时间  2019-12-20 13:29:10.6458503 +0800 CST m=+0.002985101
	fmt.Printf("%02d.%02d.%4d \n", now.Day(), now.Month(), now.Year()) //20.12.2019

	utcTime := time.Now().UTC()
	fmt.Println(now)     //2019-12-20 13:30:07.5906101 +0800 CST m=+0.002992001
	fmt.Println(utcTime) //2019-12-20 05:30:07.6055701 +0000 UTC

	testTime1()
	testTime2()
	testTime3()
}

func testTime1() {
	now := time.Now()
	addWeek := 60 * 60 * 24 * 7 * 1e9
	week_from_now := now.Add(time.Duration(addWeek))
	fmt.Println(now)           //2019-12-20 13:52:22.8047998 +0800 CST m=+0.017951501
	fmt.Println(week_from_now) //2019-12-27 13:52:22.8047998 +0800 CST m=+604800.017951501
}

func testTime2() {
	now := time.Now()
	utc := time.Now().UTC()
	fmt.Println(now.Format(time.ANSIC))  //Fri Dec 20 13:54:40 2019
	fmt.Println(now.Format(time.RFC822)) //20 Dec 19 13:54 CST

	//贼奇葩，必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5（06年1月2号下午3点4分5秒）
	fmt.Println(now.Format("2006-01-02 15:04:05")) //2019-12-20 14:02:21
	formatTime := utc.Format("20160102")
	fmt.Println(utc, "=>", formatTime) //2019-12-20 06:02:21.003534 +0000 UTC => 201261220

}

func testTime3() {
	beforNow := time.Now()
	fmt.Println(beforNow) //2019-12-20 14:49:42.7665591 +0800 CST m=+0.016987701
	time.Sleep(time.Duration(5 * 1e9))
	afternow := time.Now()
	fmt.Println(afternow) //2019-12-20 14:49:47.7669061 +0800 CST m=+5.017334701
}
