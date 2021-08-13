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

func (s *tcpClosed) Transmit(o *types.TcpOctetStream) {
	log.Println("Warning: transmitting using a closed connection.")
}

func (s *tcpClosed) ActiveOpen() TcpState {
	return NewTcpEstablished()
}

func (s *tcpClosed) PassiveOpen() TcpState {
	return NewTcpListen()
}

func (s *tcpClosed) Close() TcpState {
	return nil

}

func (s *tcpClosed) Send() TcpState {
	return nil

}

func (s *tcpClosed) Acknowledge() TcpState {
	return nil

}

func (s *tcpClosed) Synchronize() TcpState {
	return nil

}
