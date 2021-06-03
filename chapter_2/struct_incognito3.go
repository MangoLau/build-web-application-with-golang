package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Employee struct {
	Human      // 匿名字段Human
	speciality string
	phone      string // 雇员的phone字段
}

func main() {
	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-333"}
	fmt.Println("Bob's work phone is:", Bob.phone)
	// 如果我们要访问Human的phone字段
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}
