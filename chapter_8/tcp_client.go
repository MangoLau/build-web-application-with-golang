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
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkerror.CheckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkerror.CheckError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkerror.CheckError(err)
	//result, err := ioutil.ReadAll(conn)
	result := make([]byte, 256)
	_, err = conn.Read(result)
	checkerror.CheckError(err)
	fmt.Println(string(result))
	os.Exit(0)
}
