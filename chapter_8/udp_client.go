package main

import (
	"fmt"
	"github.com/mangolau/build-web-application-with-golang/chapter_8/checkerror"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		if _, err := fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0]); err != nil {
			return
		}
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkerror.CheckError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkerror.CheckError(err)
	_, err = conn.Write([]byte("anything"))
	checkerror.CheckError(err)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkerror.CheckError(err)
	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}
