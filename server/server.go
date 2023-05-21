package server

import (
	C "github.com/savion1024/wall/config"
	"github.com/savion1024/wall/tunnel"

	"log"
	"sync"
)

var (
	tcpQueue = make(chan *tunnel.ConnContext, 200)
)

func init() {
	go func() {
		queue := tcpQueue
		for c := range queue {
			go tunnel.ProcessConn(c)
		}
	}()
}

type Server struct {
	mu     sync.Mutex
	config *C.GlobalConfig
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
