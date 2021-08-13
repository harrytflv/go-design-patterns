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

func (s *tcpListen) Transmit(o *types.TcpOctetStream) {

}

func (s *tcpListen) ActiveOpen() TcpState {
	return nil
}

func (s *tcpListen) PassiveOpen() TcpState {
	return nil

}

func (s *tcpListen) Close() TcpState {
	return nil

}

func (s *tcpListen) Send() TcpState {
	return NewTcpEstablished()
}

func (s *tcpListen) Acknowledge() TcpState {
	return nil

}

func (s *tcpListen) Synchronize() TcpState {
	return nil

}
