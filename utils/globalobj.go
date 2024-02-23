package utils

import (
	"encoding/json"
	"os"

	"github.com/SakBlues/zinx/ziface"
)

// use ConfigFile to set the config. Notice that,
// the base dir is the dir where you type the command line.
const config = "conf/zinx.json"

type GlobalObj struct {
	/* server */
	// TcpServer is the global server object.
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string

	/* zinx */
	// zinx version
	Version string
	// MaxPacketSize is the max package size the server allowed.
	MaxPacketSize uint32 //都需数据包的最大值
	// MaxConn is the max connnection num server allowed.
	MaxConn int
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := os.ReadFile(config)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	// Default value.
	GlobalObject = &GlobalObj{
		Name:          "ZinxServerApp",
		Version:       "V0.4",
		TcpPort:       7777,
		Host:          "0.0.0.0",
		MaxConn:       1000,
		MaxPacketSize: 4096,
	}

	GlobalObject.Reload()
}
