package server

import "C"
import (
	global "github.com/savion1024/wall/constant"
	"net"
)

func (s *Server) StartListenHttp(in chan<- C.ConnContext) error {
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

func handleConn(conn net.Conn, in chan<- C.ConnContext) {
	// TODO find remote proxy
	// TODO exchange local and remote conn
}
