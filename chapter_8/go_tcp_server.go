package main

import (
	"fmt"
	"github.com/mangolau/build-web-application-with-golang/chapter_8/checkerror"
	"net"
	"time"
)

func main() {
	service := ":8081"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkerror.CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkerror.CheckError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	fmt.Println(conn)
	conn.Write([]byte(daytime))
}
