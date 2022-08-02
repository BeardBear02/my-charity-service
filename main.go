package main

import (
	"fmt"
	"net/http"
)

//规定函数必须接受两个值，ResponseWriter是返回网页的值，Request是接受请求。
func sayHello(w http.ResponseWriter, r *http.Request) {
	//Fprintln方法有两个返回值，不想被标记的话就用两个下划线接收一下
	_, _ = fmt.Fprintln(w, "<h1>Hello Zephyr<h1>")
	_, _ = fmt.Fprintln(w, "Hello Zephyr")
}

//想要启动的话，需要在终端 go run main.go
//启动时没有消息提示，但是确实启动了，通过localhost:8080访问
func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("http server start failure, err:%v\n", err)
		return
	}

}
