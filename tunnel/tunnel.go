package tunnel

import (
	"context"
	"io"
	"log"
	"time"
)

func ProcessConn(c *ConnContext) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*180)
	defer cancel()
	proxyAddress := c.RemoteConn.RemoteAddr().String()
	remoteConn, err := c.DialRemote(ctx, proxyAddress)
	if err != nil {
		log.Fatalf("handle remotre")
	}
	c.RemoteConn = remoteConn
	go c.exchangeConn()

}

func (c *ConnContext) exchangeConn() {
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
