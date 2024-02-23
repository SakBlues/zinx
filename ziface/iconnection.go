package ziface

import "net"

type IConnection interface {
	// Start start the read and write businses.
	Start()

	// Stop stop the connection, clean the resources.
	Stop()

	// GetTCPConnection return TCPConn, because zinx is based on TCP.
	GetTCPConnection() *net.TCPConn

	// GetConnID return connID, which uniquely identifies a connection.
	// The upper layer can use the ConnID to manage the connection.
	GetConnID() uint32

	// RemoteAddr returns the remote network address.
	RemoteAddr() net.Addr

	// TODO
	Send(data []byte) error
}

// HandFunc is a function to handle business,
// which is binded with connection.
// The function is provide by the framework caller.
type HandleFunc func(*net.TCPConn, []byte, int) error
