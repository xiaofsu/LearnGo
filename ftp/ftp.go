package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("监听 E 盘下的所有文件，监听端口为：8080")
	http.Handle("/", http.FileServer(http.Dir("E://")))
	http.ListenAndServe(":8080", nil)
}
