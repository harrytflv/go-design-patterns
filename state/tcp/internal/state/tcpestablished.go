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

func (s *tcpEstablished) Transmit(o *types.TcpOctetStream) {
	log.Println("Data transmitted.")
}

func (s *tcpEstablished) ActiveOpen() TcpState {
	return nil
}

func (s *tcpEstablished) PassiveOpen() TcpState {
	return nil

}

func (s *tcpEstablished) Close() TcpState {
	return NewTcpClosed()
}

func (s *tcpEstablished) Send() TcpState {
	return nil

}

func (s *tcpEstablished) Acknowledge() TcpState {
	return nil

}

func (s *tcpEstablished) Synchronize() TcpState {
	return nil

}
