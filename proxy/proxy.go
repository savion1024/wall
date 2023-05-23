package proxy

import (
	"github.com/savion1024/wall/common"
)

type Proxy struct {
	Host      string
	Port      int
	Name      string
	Alive     bool
	ProxyType string
	Password  string
	Sni       string
}

func (p *Proxy) LastDelay() uint16 {
	return 0
}

func (p *Proxy) Address() string {
	return common.GenAddr(p.Host, p.Port)
}
