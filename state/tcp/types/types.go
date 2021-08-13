package types

type TcpConnection interface {
	ActiveOpen() TcpConnection
	PassiveOpen() TcpConnection
	Close() TcpConnection
	Send() TcpConnection
	Acknowledge() TcpConnection
	Synchronize() TcpConnection
	ProcessOctet(o *TcpOctetStream) TcpConnection
}

type TcpOctetStream struct {
}
