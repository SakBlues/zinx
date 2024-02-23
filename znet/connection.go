package znet

import (
	"fmt"
	"net"

	"github.com/SakBlues/zinx/utils"
	"github.com/SakBlues/zinx/ziface"
)

var _ ziface.IConnection = (*Connection)(nil)

type Connection struct {
	conn *net.TCPConn

	connID uint32

	isClosed bool

	// used to call connection to exit
	ExitChan chan struct{}

	// Router contains the HandlerFunc for the IRequest.
	Router ziface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		conn:     conn,
		connID:   connID,
		isClosed: false,
		ExitChan: make(chan struct{}),
		Router:   router,
	}
	return c
}

// Read Business for Connection.
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID:", c.connID, "Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// Read client data to buf.
		buf := make([]byte, utils.GlobalObject.MaxPacketSize)
		cnt, err := c.conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err:", err)
			c.ExitChan <- struct{}{}
			continue
		}

		// Call the handleFunc bounded to connection.
		// if err := c.HandleAPI(c.conn, buf, cnt); err != nil {
		// 	fmt.Println("connID:", c.connID, ", handle err:", err)
		// 	c.ExitChan <- struct{}{}
		// 	return
		// }

		req := Request{
			conn: c,
			data: buf[:cnt],
		}

		// Call the Router to handle business.
		// TODO: error handle.
		go func(req *Request) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(&req)

		// if err := c.Router.Handle(c.conn, buf, cnt); err != nil {
		// 	fmt.Println("connID:", c.connID, ", handle err:", err)
		// 	c.ExitChan <- struct{}{}
		// 	return
		// }
	}
}

func (c *Connection) Start() {
	fmt.Println("Connection Start(), connID:", c.connID)

	go c.StartReader()

	<-c.ExitChan
	// for {
	// 	select {
	// 	case <-c.ExitChan:
	// 		// Current goroutine exit, and lead to relative goroutine exit.
	// 		return
	// 	}
	// }
}

func (c *Connection) Stop() {
	fmt.Println("Connection Stop(), connID:", c.connID)

	if c.isClosed {
		return
	}
	c.isClosed = true

	// Close conn.
	c.conn.Close()

	// Recycle resource.
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.conn
}

func (c *Connection) GetConnID() uint32 {
	return c.connID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
