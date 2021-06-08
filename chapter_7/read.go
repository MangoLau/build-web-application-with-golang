package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "mangolau.txt"
	fl, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer func(fl *os.File) {
		err := fl.Close()
		if err != nil {
			fmt.Println(userFile, err)
		}
	}(fl)
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if 0 == n {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
