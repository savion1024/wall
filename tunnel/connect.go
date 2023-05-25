package tunnel

import "C"
import (
	"context"
	global "github.com/savion1024/wall/constant"
	"io"
	"net"
	"time"

	"github.com/gofrs/uuid/v5"
)

type ConnContext struct {
	ConnId     uuid.UUID
	WorkMode   global.WorkMode
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

func NewConnContext(w global.WorkMode) *ConnContext {
	id, _ := uuid.NewV4()
	return &ConnContext{ConnId: id, WorkMode: w}
}

func (c *ConnContext) exchangeConnData() {
	ch := make(chan error)
	go func() {
		_, err := io.Copy(WriteOnlyWriter{Writer: c.LocalConn}, ReadOnlyReader{Reader: c.RemoteConn})
		c.LocalConn.SetReadDeadline(time.Now())
		ch <- err
	}()
	io.Copy(WriteOnlyWriter{Writer: c.RemoteConn}, ReadOnlyReader{Reader: c.LocalConn})
	c.RemoteConn.SetReadDeadline(time.Now())
	<-ch
}
