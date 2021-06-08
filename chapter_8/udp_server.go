package main

import (
	"github.com/mangolau/build-web-application-with-golang/chapter_8/checkerror"
	"net"
	"time"
)

func main() {
	service := ":8080"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkerror.CheckError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkerror.CheckError(err)
	for {
		udpHandleClient(conn)
	}
}

func udpHandleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
