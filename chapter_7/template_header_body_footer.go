package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
	//s1.ExecuteTemplate(os.Stdout, "header", nil)
	fmt.Println()

	// 因为嵌套了header和footer的内容，就会同时输出三个的内容。
	s1.ExecuteTemplate(os.Stdout, "content", nil)
	fmt.Println()

	//s1.ExecuteTemplate(os.Stdout, "footer", nil)
	fmt.Println()

	// 没有任何的输出，因为在默认的情况下没有默认的子模板，所以不会输出任何的东西。
	s1.Execute(os.Stdout, nil)
}
