package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("空 pipeline if demo: {{if ``}} 不会输出。{{end}}\n"))
	if err := tEmpty.Execute(os.Stdout, nil); err != nil {
		fmt.Println("excute 1 error: ", err)
		return
	}

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("不为空的 pipeline if demo: {{if `anything`}} 我有内容，我会输出！{{end}}\n"))
	if err := tWithValue.Execute(os.Stdout, nil); err != nil {
		fmt.Println("excute 2 error: ", err)
		return
	}

	// if里面无法使用条件判断，例如.Mail=="astaxie@gmail.com"，这样的判断是不正确的，if里面只能是bool值
	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} if部分{{else}}else 部分{{end}}\n"))
	if err := tIfElse.Execute(os.Stdout, nil); err != nil {
		fmt.Println("excute 2 error: ", err)
		return
	}
}
