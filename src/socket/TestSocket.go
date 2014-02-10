package socket

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	MAX_CONN_NUM = 5
)

func EchoFunc(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			return
		}
		_, err = conn.Write(buf)
		if err != nil {
			return
		}
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("ip4", "0.0.0.0:8888")
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("Error listening:", err.Error)
		os.Exit(1)
	}

	defer listener.Close()

	var cur_conn_num int = 0
	conn_chan := make(chan net.Conn)
	ch_conn_change := make(chan int)

	go func() {
		for conn_change := range ch_conn_change {
			cur_conn_num += conn_change
		}
	}()

	go func() {
		for _ = range time.Tick(1e8) {
			fmt.Println("cur conn num:%f\n", cur_conn_num)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accept:", err.Error())
			return
		}
		conn_chan <- conn
	}
}
