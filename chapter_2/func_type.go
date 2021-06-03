package main

import "fmt"

type testInt func(int) bool // 声明了一个函数类型

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd) // 函数当作值来传递了
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are； ", even)
}

func isOdd(num int) bool {
	return num%2 == 0
}

func isEven(num int) bool {
	return num%2 != 0
}

// 声明的函数类型在这个地方当做了一个参数

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
