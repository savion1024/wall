package tunnel

import (
	"context"
	"time"
)

func ProcessConn(c *ConnContext) {
	_, cancel := context.WithTimeout(context.Background(), time.Second*180)
	defer cancel()
	//proxy :=
	//remoteConn, err := c.DialRemote(ctx, proxyAddress)
	//if err != nil {
	//	log.Fatalf("handle remotre")
	//}
	//c.RemoteConn = remoteConn
	//go c.exchangeConnData()

}
