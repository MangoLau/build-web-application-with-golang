package main

import "fmt"

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic
	// 另外记住一点的就是channel不像文件之类的，不需要经常去关闭，
	// 只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的
	close(c)
}
