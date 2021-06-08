package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	UserName string
}

func main() {
	t := template.New("filename example")
	// 这个字段必须是导出的(字段首字母必须是大写的)，否则在渲染的时候就会报错
	t, _ = t.Parse("hello {{.UserName}}!")
	p := Person{UserName: "MangoLau"}
	err := t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println(err)
		return
	}
}
