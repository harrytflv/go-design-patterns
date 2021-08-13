package state

import (
	"github.com/harrytflv/go-design-patterns/state/tcp/types"
)

// For simplicity, error is not handled.
type TcpState interface {
	ActiveOpen() TcpState
	PassiveOpen() TcpState
	Close() TcpState
	Send() TcpState
	Acknowledge() TcpState
	Synchronize() TcpState
	Transmit(o *types.TcpOctetStream) TcpState
}
