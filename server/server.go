package server

import (
	"fmt"
	"github.com/savion1024/wall/logger"
	"sync"

	C "github.com/savion1024/wall/config"
	"github.com/savion1024/wall/tunnel"
)

var (
	tcpQueue = make(chan *tunnel.ConnContext, 200)
)

func init() {
	go func() {
		queue := tcpQueue
		for c := range queue {
			fmt.Println(fmt.Sprintf("catch connect: %s", c.ID()))
			go tunnel.ProcessConn(c)
		}
	}()
}

type Server struct {
	mu     sync.Mutex
	config *C.GlobalConfig
	logger *logger.Logger
}

func NewServer(g *C.GlobalConfig) (*Server, error) {
	s := &Server{
		config: g,
		logger: logger.NewStdLogger(true, false, true, true, false),
	}
	return s, nil
}

func (s *Server) Run() {
	err := s.StartListenHttp(tcpQueue)
	if err != nil {
		s.logger.Fatalf("Run server failed: %s", err.Error())
	}
}

func (s *Server) PrintBaseConfig() {
	fmt.Println("********************* Wall *********************")
	fmt.Println("***                                          ***")
	fmt.Println("***       科学来讲, 这太不科学。             ***")
	fmt.Println("***                     -《要讲科学》        ***")
	fmt.Println("***                                          ***")
	fmt.Println("***                                          ***")
	fmt.Println("************************************************")
	fmt.Println(" ")
	fmt.Println(fmt.Sprintf("   WorkMode:    %s ", s.config.WorkMode))
	fmt.Println(fmt.Sprintf("   ProxyMode:   %s ", s.config.L.ProxyMode.String()))
	fmt.Println(fmt.Sprintf("   Http Listen: %s ", s.config.L.HttpAddress()))
}
