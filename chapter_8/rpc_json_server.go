package main

import (
	"errors"
	"github.com/mangolau/build-web-application-with-golang/chapter_8/checkerror"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	if err := rpc.Register(arith); err != nil {
		log.Fatal("rpc Register error ", err)
		return
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkerror.CheckError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkerror.CheckError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)
	}
}
