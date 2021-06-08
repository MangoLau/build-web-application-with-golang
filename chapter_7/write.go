package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "mangolau.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer func(fout *os.File) {
		err := fout.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(fout)

	for i := 0; i < 10; i++ {
		if _, err := fout.WriteString("Just a test!\r\n"); err != nil {
			return
		}
		if _, err := fout.Write([]byte("Just a test!\r\n")); err != nil {
			return
		}
	}
}
