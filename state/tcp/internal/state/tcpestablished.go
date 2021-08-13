package state

import (
	"log"

	"github.com/harrytflv/go-design-patterns/state/tcp/types"
)

type tcpEstablished struct{}

func NewTcpEstablished() TcpState {
	log.Println("Connection established.")
	return &tcpEstablished{}
}

func (s *tcpEstablished) Transmit(o *types.TcpOctetStream) TcpState {
	log.Println("Data transmitted.")
	return s
}

func (s *tcpEstablished) ActiveOpen() TcpState {
	return s
}

func (s *tcpEstablished) PassiveOpen() TcpState {
	return s
}

func (s *tcpEstablished) Close() TcpState {
	return NewTcpClosed()
}

func (s *tcpEstablished) Send() TcpState {
	return s
}

func (s *tcpEstablished) Acknowledge() TcpState {
	return s
}

func (s *tcpEstablished) Synchronize() TcpState {
	return s
}
