package main

import (
	"github.com/mangolau/build-web-application-with-golang/chapter_8/checkerror"
	"net"
	"time"
)

func main() {
	service := ":8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkerror.CheckError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkerror.CheckError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}
