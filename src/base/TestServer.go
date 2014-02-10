package main

import (
	//"base/protocol"
	"base/socket"
)

var (
	socketSrv *socket.SocketServer
)

func main() { 
	config := socket.NewConfig()
	config.CloseingTimeout = 10000
	config.HeartbeatTimeout = 1000 * 30
	config.Addr = "0.0.0.0:9000"
	socketSrv = socket.NewSocketServer(config)
	socketSrv.Start()

}
