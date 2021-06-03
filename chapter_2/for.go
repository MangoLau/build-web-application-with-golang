package main

import "fmt"

func main() {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println("sum2 is equal to ", sum2)

	for index := 10; index > 0; index-- {
		if index == 5 {
			break // 或者continue
		}
		fmt.Println(index)
	}
}
