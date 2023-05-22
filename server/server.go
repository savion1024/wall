package server

import (
	"fmt"
	global "github.com/savion1024/wall/constant"
	logger2 "github.com/savion1024/wall/logger"
	"sync"

	C "github.com/savion1024/wall/config"
	"github.com/savion1024/wall/tunnel"
)

var (
	tcpQueue = make(chan *tunnel.ConnContext, 200)
	logger   *logger2.Logger
)

func init() {
	logger = logger2.NewStdLogger(true, false, true, true, false)
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
}

func NewServer(g *C.GlobalConfig) *Server {
	s := &Server{
		config: g,
	}
	return s
}

func (s *Server) Run() {
	err := s.StartListenHttp(tcpQueue)
	if err != nil {
		logger.Fatalf("Run server failed: %s", err.Error())
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
	if s.config.L.ProxyMode == global.HTTP {
		fmt.Println(fmt.Sprintf("   Http Listen: %s ", s.config.L.HttpAddress()))
	}
	fmt.Println(fmt.Sprintf("   Start the service successfully and enjoy surfing. "))
}
