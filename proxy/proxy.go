package proxy

import "net"

type Proxy struct {
	host      string
	port      string
	name      string
	alive     bool
	proxyType string
}

func (p *Proxy) Name() string {
	return p.name
}

func (p *Proxy) Alive() bool {
	return p.alive
}

func (p *Proxy) LastDelay() uint16 {
	return 0
}

func (p *Proxy) Address() string {
	return net.JoinHostPort(p.host, p.port)
}
