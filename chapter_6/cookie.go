package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", sayhelloName)       // 设置访问的路由
	err := http.ListenAndServe(":8080", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Scheme)
	fmt.Println(r.Form["url_log"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello MangoLau!\n") // 这个写入到 w 的是输出到客户端的

	// cookie 设置
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "mangolau", Expires: expiration}
	http.SetCookie(w, &cookie)

	// 读取cookie
	// cookieV, _ := r.Cookie("username")
	// fmt.Fprintf(w, cookieV)
	for _, cookieItem := range r.Cookies() {
		fmt.Fprintf(w, cookieItem.Name)
		fmt.Fprintln(w)
	}
}
