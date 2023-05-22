package tunnel

import (
	"context"
	"log"
	"time"
)

func ProcessConn(c *ConnContext) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*180)
	defer cancel()
	proxyAddress := "hk.aliyun-cave.xyz:41501"
	remoteConn, err := c.DialRemote(ctx, proxyAddress)
	if err != nil {
		log.Fatalf(err.Error())
	}
	c.RemoteConn = remoteConn
	go c.exchangeConnData()

}
