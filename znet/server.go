package znet

import (
	"fmt"
	"net"

	"github.com/SakBlues/zinx/ziface"
)

// Server is based on TCP.
type Server struct {
	name      string
	ipVersion string
	ip        string
	port      int
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, Starting...\n", s.ip, s.port)

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

			// TODO
			// 3.3 Handler business. Handler and conn should be bound at this point.
			// Here is a demo: an echo service of up to 512 bytes.
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Receive buf err:", err)
						continue
					}

					fmt.Printf("Receive client buf: %s, cnt = %d\n", buf, cnt)

					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("Write back buf err:", err)
						continue
					}
				}
			}()
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

func NewServer(name string) ziface.IServer {
	s := &Server{
		name:      name,
		ipVersion: "tcp4",
		ip:        "0.0.0.0",
		port:      8999,
	}

	return s
}
