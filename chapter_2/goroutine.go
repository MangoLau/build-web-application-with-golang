package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() // 表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。
		fmt.Println(s)
	}
}

func main() {
	runtime.GOMAXPROCS(4)
	go say("world") // 开一个新的 Goroutines 执行
	say("hello")    // 当前 Goroutines 执行
}
