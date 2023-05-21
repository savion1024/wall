package tunnel

import "C"
import (
	"context"
	"github.com/gofrs/uuid/v5"
	"io"
	"net"
	"time"
)

type ConnContext struct {
	ConnId     uuid.UUID
	LocalConn  net.Conn
	RemoteConn net.Conn
}

func (c *ConnContext) ID() uuid.UUID {
	return c.ConnId
}

func (c *ConnContext) DialRemote(ctx context.Context, address string) (net.Conn, error) {
	dialer := &net.Dialer{}
	remoteConn, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return nil, err
	}
	return remoteConn, nil
}

func NewConnContext() *ConnContext {
	id, _ := uuid.NewV4()
	return &ConnContext{ConnId: id}
}

func (c *ConnContext) exchangeConnData() {
	ch := make(chan error)
	leftConn := c.LocalConn
	rightConn := c.RemoteConn
	go func() {
		_, err := io.Copy(WriteOnlyWriter{Writer: leftConn}, ReadOnlyReader{Reader: rightConn})
		leftConn.SetReadDeadline(time.Now())
		ch <- err
	}()

	io.Copy(WriteOnlyWriter{Writer: rightConn}, ReadOnlyReader{Reader: leftConn})
	rightConn.SetReadDeadline(time.Now())
	<-ch
}
