package server

import (
	C "github.com/savion1024/wall/config"
	"log"
	"net"
	"sync"
)

var (
	tcpQueue = make(chan ConnContext, 200)
)

type Server struct {
	mu     sync.Mutex
	config *C.GlobalConfig
	conn   net.Conn
}

func NewServer(g *C.GlobalConfig) (*Server, error) {
	s := &Server{
		config: g,
	}
	return s, nil
}

func (s *Server) Run() {
	err := s.StartListenHttp(tcpQueue)
	if err != nil {
		log.Fatalf("start http error")
	}

}
