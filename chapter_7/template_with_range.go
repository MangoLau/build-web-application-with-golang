package main

import (
	"fmt"
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "xushiwei"}
	t := template.New("fieldname example")
	/*
		{{range}} 这个和Go语法里面的range类似，循环操作数据
		{{with}}操作是指当前对象的值，类似上下文的概念
	*/
	t, _ = t.Parse(`hello {{.UserName}}!
		{{range .Emails}}
			an email {{.}}
		{{end}}
		{{with .Friends}}
		{{range .}}
			my friend name is {{.Fname}}
		{{end}}
		{{end}}
	`)
	p := Person{UserName: "MangoLau",
		Emails:  []string{"1@q.com", "2@q.com"},
		Friends: []*Friend{&f1, &f2}}
	err := t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println(err)
		return
	}
}
