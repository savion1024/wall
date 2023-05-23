package config

import (
	"errors"
	"github.com/savion1024/wall/common"
	"github.com/savion1024/wall/proxy"
	"gopkg.in/yaml.v3"
	"os"

	C "github.com/savion1024/wall/constant"
)

type GlobalConfig struct {
	L        *LocalConfig
	Proxies  map[string]Proxy `json:"proxies"`
	Rules    []C.Rules
	WorkMode C.WorkMode `json:"work-mode"`
}

type Proxy interface {
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
		return nil, errors.New("need config file")
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
		L: &LocalConfig{
			HttpProxyPort: C.DefaultPort,
		},
		Proxies:  map[string]Proxy{},
		WorkMode: raw.WorkMode,
	}
	if raw.HttpPort != 0 {
		g.L.HttpProxyPort = raw.HttpPort
	}
	if raw.MixedPort != 0 {
		g.L.MixedProxyPort = raw.MixedPort
	}
	if raw.SocksPort != 0 {
		g.L.SocksProxyPort = raw.SocksPort
	}
	// parse proxies
	for _, p := range raw.Proxy {
		op := &proxy.Proxy{}
		op.Name = p["name"].(string)
		op.Host = p["server"].(string)
		op.Port = p["port"].(int)
		op.ProxyType = p["type"].(string)
		op.Password = p["password"].(string)
		op.Sni = p["sni"].(string)
		g.Proxies[op.Name] = op
	}
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
