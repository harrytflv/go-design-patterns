package main

import (
	"log"

	"github.com/harrytflv/go-design-patterns/state/tcp/tcpconnection"
	"github.com/harrytflv/go-design-patterns/state/tcp/types"
)

func main() {
	log.Println("Run1")
	tcpconnection.NewTcpConnection().ActiveOpen().ProcessOctet(&types.TcpOctetStream{}).Close()
	log.Println("Run2")
	tcpconnection.NewTcpConnection().PassiveOpen().Send().ProcessOctet(&types.TcpOctetStream{}).Close()
}
