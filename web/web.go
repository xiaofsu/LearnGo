package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Rope string

func helloHandler(writer http.ResponseWriter, req *http.Request) {
	var a Rope = "111"
	fmt.Println("----> handler start <----" + a)
	fmt.Println("This is a raw string \n")
	//writer.Header().Set("content-type","application/json")
	fmt.Fprintf(writer, "hello go, rquest url is %s", req.URL.Path)
}
func main() {
	fmt.Println("----> go start <----")
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/test", test)
	http.HandleFunc("/index", index)
	err := http.ListenAndServe("localhost:9090", nil)
	checkErr(err)
}
func test(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("----> test start <----")
	fmt.Fprint(writer, "这个GoLang蛮厉害的")
}
func index(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("----> index handler start <----")
	t, err := template.ParseFiles("template/index.html")
	checkErr(err)
	t.Execute(writer, nil)
}
func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
