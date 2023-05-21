package tunnel

import "C"
import (
	"context"
	"github.com/gofrs/uuid/v5"
	"net"
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
