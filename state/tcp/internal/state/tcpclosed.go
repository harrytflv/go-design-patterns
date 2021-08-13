package state

import (
	"log"

	"github.com/harrytflv/go-design-patterns/state/tcp/types"
)

type tcpClosed struct{}

func NewTcpClosed() TcpState {
	log.Println("Connection closed.")
	return &tcpClosed{}
}

func (s *tcpClosed) Transmit(o *types.TcpOctetStream) TcpState {
	log.Println("Warning: transmitting using a closed connection.")
	return s
}

func (s *tcpClosed) ActiveOpen() TcpState {
	return NewTcpEstablished()
}

func (s *tcpClosed) PassiveOpen() TcpState {
	return NewTcpListen()
}

func (s *tcpClosed) Close() TcpState {
	return s
}

func (s *tcpClosed) Send() TcpState {
	return s

}

func (s *tcpClosed) Acknowledge() TcpState {
	return s

}

func (s *tcpClosed) Synchronize() TcpState {
	return s
}
