package znet

import (
	"fmt"
	"net"

	"github.com/SakBlues/zinx/ziface"
)

var _ ziface.IConnection = (*Connection)(nil)

type Connection struct {
	conn *net.TCPConn

	connID uint32

	isClosed bool

	handleAPI ziface.HandFunc

	// used to call connection to exit
	exitChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandFunc) *Connection {
	c := &Connection{
		conn:      conn,
		connID:    connID,
		isClosed:  false,
		handleAPI: callback_api,
		exitChan:  make(chan bool),
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
		buf := make([]byte, 512)
		cnt, err := c.conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err:", err)
			c.exitChan <- true
			continue
		}

		// Call the handleFunc bounded to connection.
		if err := c.handleAPI(c.conn, buf, cnt); err != nil {
			fmt.Println("connID:", c.connID, ", handle err:", err)
			c.exitChan <- true
			return
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Connection Start(), connID:", c.connID)

	go c.StartReader()

	for {
		select {
		case <-c.exitChan:
			// Current goroutine exit, and lead to relative goroutine exit.
			return
		}
	}
}

func (c *Connection) Stop() {
	fmt.Println("Connection Stop(), connID:", c.connID)

	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// Close conn.
	c.conn.Close()

	// Recycle resource.
	close(c.exitChan)
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
