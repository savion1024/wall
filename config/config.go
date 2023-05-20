package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"net"
	"os"

	C "github.com/savion1024/wall/constant"
)

type GlobalConfig struct {
	LC      *LocalConfig
	Proxies map[string]*Proxy
}

type Proxy interface {
	Name() string
	Alive() bool
	LastDelay() uint16
	Dial() (net.Conn, error)
}

type LocalConfig struct {
	ProxyMode      C.WorkMode
	HttpProxyAddr  string
	HttpProxyPort  int
	MixedProxyAddr string
	MixedProxyPort int
	SocksProxyAddr string
	SocksProxyPort int
}

// Parse config from []byte
func Parse(filePtah string) (*GlobalConfig, error) {
	if filePtah == "" {
		return nil, errors.New("i need yaml file so can start me")
	}
	if _, err := os.Stat(filePtah); err != nil {
		return nil, errors.New("can't find file")
	}
	data, err := os.ReadFile(filePtah)
	if err != nil {
		return nil, err
	}
	g := &GlobalConfig{LC: &LocalConfig{
		HttpProxyPort: C.DefaultPort,
	}}
	if err := yaml.Unmarshal(data, g); err != nil {
		return nil, err
	}
	// TODO check config
	return g, nil
}
