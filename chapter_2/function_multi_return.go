package main

import "fmt"

func SumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

// SumAndProduct2 最好命名返回值，因为不命名返回值，虽然使得代码更加简洁了，但是会造成生成的文档可读性差
func SumAndProduct2(a, b int) (sum, product int) {
	sum, product = a+b, a*b
	return
}

func main() {
	x, y := 3, 4
	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)

	xPLUSy2, xTIMESy2 := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy2)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy2)
}
