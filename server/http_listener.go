package server

import (
	global "github.com/savion1024/wall/constant"
	"github.com/savion1024/wall/tunnel"
	"net"
)

// StartListenHttp catch http request
func (s *Server) StartListenHttp(in chan<- *tunnel.ConnContext) error {
	if s.config.LC.ProxyMode != global.HTTP {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	addr := s.config.LC.HttpProxyAddr
	port := s.config.LC.HttpProxyPort
	bindAddress := net.JoinHostPort(addr, port)
	l, err := net.Listen("tcp", bindAddress)
	if err != nil {
		return err
	}
	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				continue
			}
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
	in <- c

}
