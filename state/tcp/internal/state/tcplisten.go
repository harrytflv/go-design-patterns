package state

import (
	"log"

	"github.com/harrytflv/go-design-patterns/state/tcp/types"
)

type tcpListen struct{}

func NewTcpListen() TcpState {
	log.Println("Connection listening.")
	return &tcpListen{}
}

func (s *tcpListen) Transmit(o *types.TcpOctetStream) TcpState {
	return s
}

func (s *tcpListen) ActiveOpen() TcpState {
	return s
}

func (s *tcpListen) PassiveOpen() TcpState {
	return s
}

func (s *tcpListen) Close() TcpState {
	return s
}

func (s *tcpListen) Send() TcpState {
	return NewTcpEstablished()
}

func (s *tcpListen) Acknowledge() TcpState {
	return s
}

func (s *tcpListen) Synchronize() TcpState {
	return s
}
