package znet

import (
	"fmt"
	"net"

	"github.com/SakBlues/zinx/utils"
	"github.com/SakBlues/zinx/ziface"
)

var _ ziface.IServer = (*Server)(nil)

// Server is based on TCP.
type Server struct {
	name      string
	ipVersion string
	ip        string
	port      int

	Router ziface.IRouter
}

// TODO: Remove.
// func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
// 	fmt.Println("[Conn Handle] CallBackToClient ... ")
// 	if _, err := conn.Write(data[:cnt]); err != nil {
// 		fmt.Println("Write back buf err:", err)
// 		return errors.New("CallBackToClient error")
// 	}
// 	return nil
// }

func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s,listenner at ip: %s, port %d is starting\n",
		s.name, s.ip, s.port)
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d,  MaxPacketSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)

	// Listener Business
	go func() {
		//  Here is a go TCP server development process.

		// 1. Get address.
		addr, err := net.ResolveTCPAddr(s.ipVersion, fmt.Sprintf("%s:%d", s.ip, s.port))
		if err != nil {
			fmt.Println("Resolve tcp addr err:", err)
			return
		}

		// 2. Listen server address.
		listenner, err := net.ListenTCP(s.ipVersion, addr)
		if err != nil {
			fmt.Println("Listen:", s.ipVersion, ", err:", err)
			return
		}

		fmt.Println("Start Zinx server:", s.name, "success. Listenning...")

		cid := uint32(0)
		// 3. Accept and handle business.
		for {
			//3.1 Block to wait for client connect.
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err:", err)
				continue
			}

			// TODO
			// 3.2 Validate config, e.g., close the connection if exceed the max connections, etc.

			// 3.3 Handler business. Handler and conn should be bound at this point.
			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server:", s.name)

	//TODO  Server.Stop()
	// Stop or recycle server resources.
}

func (s *Server) Serve() {
	s.Start()

	//TODO Server.Serve()
	// Do some extra work.

	// Block to avoid listenner go exit.
	select {}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add router success.")
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		name:      utils.GlobalObject.Name,
		ipVersion: "tcp4",
		ip:        utils.GlobalObject.Host,
		port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
	return s
}
