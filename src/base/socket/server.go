package socket

import (
	"base/protocol"
	//"fmt"
	"log"
	"net"
	"sync"
	"time"
)

var (
	tcp string = "tcp"
	ip4 string = "ip4"
)

const (
	Heartbeat_timeout     = 15000     //心跳超时
	Closing_timeout       = 10000     // 关闭超时
	Session_maintain_time = 1000 * 60 //session的维护周期
	Listening_port        = 8888      //默认监听的端口
)

type Config struct {
	HeartbeatTimeout int
	CloseingTimeout  int
	Addr             string
	CheckSessionTime int //检查session
	MessageHandler   func(protoPack *protocol.ProtoPack)
}

func NewConfig() *Config {
	return &Config{}
}

type SocketSession struct {
	conn              *net.Conn
	connectedTime     int64
	lastHeartbeatTime int64
}

type SocketServer struct {
	heartbeatTimeout    int
	closingTimeout      int
	sessionMaintainTime int
	addr                string
	mutex               sync.RWMutex
	sessions            map[int]*SocketSession
	listener            *net.TCPListener
	messageHandler      func(protoPack *protocol.ProtoPack)
}

func NewSocketServer(config *Config) *SocketServer {
	server := &SocketServer{}
	if config != nil {
		server.closingTimeout = config.CloseingTimeout
		server.heartbeatTimeout = config.HeartbeatTimeout
		server.addr = config.Addr
		server.messageHandler = config.MessageHandler
		server.sessionMaintainTime = config.CheckSessionTime
	}

	if server.heartbeatTimeout == 0 {
		server.heartbeatTimeout = Heartbeat_timeout
	}

	if server.closingTimeout == 0 {
		server.closingTimeout = Closing_timeout
	}

	server.sessions = make(map[int]*SocketSession)
	return server
}

/**
 * 启动服务器
 * @author abram
 */
func (this *SocketServer) Start() error {
	tcpAddr, err := net.ResolveTCPAddr(tcp, this.addr)
	if err != nil {
		log.Fatal("服务器监听IP设置错误，信息：" + err.Error())
		return err
	}

	listener, err := net.ListenTCP(tcp, tcpAddr)

	if err != nil {
		log.Fatal("启动服务器错误，信息：" + err.Error())
		return err
	}
	this.listener = listener
	defer listener.Close()
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Println("客户端接入错误，信息：" + err.Error())
				continue
			}

			go this.connectionHandler(conn)
		}
	}()

	go this.sessionMaintainHandler(this.sessionMaintainTime)
	log.Println("===服务器启动成功。")
	return nil
}

/**
 * 客户端接入管理
 * @author abram
 */
func (this *SocketServer) connectionHandler(conn net.Conn) {

}

/**
 * 添加session
 * @author abram
 * @param key 一般使用userId
 * @param session SocketSession
 */
func (this *SocketServer) AddSession(key int, session *SocketSession) {

}

/**
 * 移除session
 * @author abram
 * @param key 一般使用userId
 * @return SocketSession 当指定的SocketSession 不存在时返回nil
 */
func (this *SocketServer) RemoveSession(key int) *SocketSession {
	return nil
}

/**
 * 获取session长度
 * @author abram
 * @return int
 */
func (this *SocketServer) GetSessionSize() int {
	return 0
}

/**
 * 关闭服务
 * @author abram
 */
func (this *SocketServer) Close() {

}

/**
 * 维护session的协程
 * @author abram
 * @param time 检查周期
 */
func (this *SocketServe) sessionMaintainHandler(time int) {

}
