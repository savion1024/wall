package server

import (
	"fmt"
	"net"

	"github.com/savion1024/wall/common"
	global "github.com/savion1024/wall/constant"
	"github.com/savion1024/wall/tunnel"
)

// StartListenHttp catch http request
func (s *Server) StartListenHttp(in chan<- *tunnel.ConnContext) error {
	if s.config.L.ProxyMode != global.HTTP {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	port := s.config.L.HttpProxyPort
	l, err := net.Listen("tcp", common.GenAddr("*", port))
	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				continue
			}
			fmt.Println("new connect accept")
			go handleConn(conn, in)
		}
	}()
	return nil
}

// handleConn handle local http connect
func handleConn(conn net.Conn, in chan<- *tunnel.ConnContext) {
	// TODO new connContext and push in tcp queue
	c := tunnel.NewConnContext()
	c.LocalConn = conn
	fmt.Println(fmt.Sprintf("connect id : %s", c.ID()))
	in <- c
}
