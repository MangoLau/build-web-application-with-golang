package main

import (
	"fmt"
	"os"
	"regexp"
)

func IsIP(ip string) (m bool) {
	m, _ = regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}$", ip)
	return
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Usage: regexp [string]")
	}
	if m, _ := regexp.MatchString("^[0-9]+$", os.Args[1]); m {
		fmt.Println("数字")
	} else {
		fmt.Println("不是数字")
	}

	if IsIP("122.22.22.1") {
		fmt.Println("是合法的IP地址")
	} else {
		fmt.Println("不是合法的IP地址")
	}
}
