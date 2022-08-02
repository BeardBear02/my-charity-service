package main

import (
	"fmt"
	"net/http"
)

//规定函数必须接受两个值，ResponseWriter是返回网页的值，Request是接受请求。
func sayHello(w http.ResponseWriter, r *http.Request) {
	//Fprintln方法有两个返回值，不想被标记的话就用两个下划线接收一下
	_, _ = fmt.Fprintln(w, "Hello Zephyr")
}
func main() {
	http.HandleFunc("/hello", sayHello)

}
