package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("监听 C 盘下的所有文件，监听端口为：9999")
	http.Handle("/", http.FileServer(http.Dir("C://MapBox-master")))
	http.ListenAndServe(":80", nil)
}
