package config

import (
	"net"
)

type GlobalConfig struct {
	LocalConfig
	Proxies map[string]*Proxy
}

type Proxy interface {
	Name() string
	Alive() bool
	LastDelay() uint16
	Dial() (net.Conn, error)
}

type LocalConfig struct {
	ProxyMode      WorkMode
	HttpProxyAddr  string
	HttpProxyPort  int
	MixedProxyAddr string
	MixedProxyPort int
	SocksProxyAddr string
	SocksProxyPort int
}
