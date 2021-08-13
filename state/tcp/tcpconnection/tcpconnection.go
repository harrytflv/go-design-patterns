package tcpconnection

import (
	"log"

	"github.com/harrytflv/go-design-patterns/state/tcp/internal/state"
	"github.com/harrytflv/go-design-patterns/state/tcp/types"
)

type tcpConnectionImpl struct {
	state state.TcpState
}

func NewTcpConnection() types.TcpConnection {
	log.Println("Creating a new connection.")
	return &tcpConnectionImpl{
		state: state.NewTcpClosed(),
	}
}

func (t *tcpConnectionImpl) ActiveOpen() types.TcpConnection {
	return &tcpConnectionImpl{
		state: t.state.ActiveOpen(),
	}
}

func (t *tcpConnectionImpl) PassiveOpen() types.TcpConnection {
	return &tcpConnectionImpl{
		state: t.state.PassiveOpen(),
	}
}

func (t *tcpConnectionImpl) Close() types.TcpConnection {
	return &tcpConnectionImpl{
		state: t.state.Close(),
	}
}

func (t *tcpConnectionImpl) Send() types.TcpConnection {
	return &tcpConnectionImpl{
		state: t.state.Send(),
	}
}

func (t *tcpConnectionImpl) Acknowledge() types.TcpConnection {
	return &tcpConnectionImpl{
		state: t.state.Acknowledge(),
	}
}

func (t *tcpConnectionImpl) Synchronize() types.TcpConnection {
	return &tcpConnectionImpl{
		state: t.state.Synchronize(),
	}
}

func (t *tcpConnectionImpl) ProcessOctet(o *types.TcpOctetStream) types.TcpConnection {
	t.state.Transmit(o)
	return t
}
