package znet

import "github.com/SakBlues/zinx/ziface"

var _ ziface.IRequest = (*Request)(nil)

type Request struct {
	conn ziface.IConnection
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
