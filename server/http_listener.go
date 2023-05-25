package server

import (
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
				logger.Errorf("Handle conn accept error: %s", err.Error())
			}
			go handleConn(s.config.WorkMode, conn, in)
		}
	}()
	return nil
}

// handleConn handle local http connect
func handleConn(w global.WorkMode, conn net.Conn, in chan<- *tunnel.ConnContext) {
	// TODO new connContext and push in tcp queue
	c := tunnel.NewConnContext(w)
	c.LocalConn = conn
	in <- c
}
