package main

import (
	"fmt"
	"net"
	"time"

	"github.com/SakBlues/zinx/utils"
)

// Mock client
func main() {

	fmt.Println("Client start ...")
	// wait server to start.
	time.Sleep(1 * time.Second)

	// 1. Connect to server.
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", utils.GlobalObject.Host,
		utils.GlobalObject.TcpPort))
	if err != nil {
		fmt.Println("Client start failed, exit!")
		return
	}

	// 2. Write data.
	for {
		_, err := conn.Write([]byte("Hello Zinx V0.4."))
		if err != nil {
			fmt.Println("Write conn err: ", err)
			return
		}

		buf := make([]byte, utils.GlobalObject.MaxPacketSize)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err: ", err)
			return
		}

		fmt.Printf("Server call back: %s, cnt = %d\n", buf, cnt)

		// cpu block
		time.Sleep(1 * time.Second)
	}
}
