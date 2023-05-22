package config

import (
	"errors"
	"github.com/savion1024/wall/common"
	"gopkg.in/yaml.v3"
	"os"

	C "github.com/savion1024/wall/constant"
)

type GlobalConfig struct {
	L        *LocalConfig
	Proxies  map[string]*Proxy `json:"proxies"`
	WorkMode C.WorkMode        `json:"work-mode"`
}

type Proxy interface {
	Name() string
	Alive() bool
	LastDelay() uint16
	Address() string
}

type LocalConfig struct {
	ProxyMode      C.ProxyMode `json:"proxy-mode"`
	HttpProxyPort  int         `json:"http-proxy-port"`
	MixedProxyPort int         `json:"mixed-proxy-port"`
	SocksProxyPort int         `json:"socks-proxy-port"`
}

func (l *LocalConfig) HttpAddress() string {
	return common.GenAddr("*", l.HttpProxyPort)
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
	raw := &RawConfig{}
	if err := yaml.Unmarshal(data, raw); err != nil {
		return nil, err
	}
	g := ParseRawConfig(raw)
	// TODO check config
	return g, nil
}

func ParseRawConfig(raw *RawConfig) *GlobalConfig {
	g := &GlobalConfig{
		L:        &LocalConfig{},
		Proxies:  map[string]*Proxy{},
		WorkMode: raw.WorkMode,
	}
	g.L.HttpProxyPort = raw.HttpPort
	g.L.MixedProxyPort = raw.MixedPort
	g.L.SocksProxyPort = raw.SocksPort
	return g

}

type RawConfig struct {
	HttpPort  int        `yaml:"http-port"`
	SocksPort int        `yaml:"socks-port"`
	MixedPort int        `yaml:"mixed-port"`
	AllowLan  bool       `yaml:"allow-lan"`
	WorkMode  C.WorkMode `yaml:"mode"`

	Proxy []map[string]any `yaml:"proxies"`
}
