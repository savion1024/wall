package config

import (
	"errors"
	"fmt"
	"github.com/savion1024/wall/common"
	"github.com/savion1024/wall/proxy"
	"github.com/savion1024/wall/rules"
	"gopkg.in/yaml.v3"
	"os"
	"strings"

	C "github.com/savion1024/wall/constant"
)

type GlobalConfig struct {
	L        *LocalConfig
	Proxies  map[string]Proxy `json:"proxies"`
	Rules    []rules.Rule
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

type RawConfig struct {
	HttpPort  int        `yaml:"http-port"`
	SocksPort int        `yaml:"socks-port"`
	MixedPort int        `yaml:"mixed-port"`
	AllowLan  bool       `yaml:"allow-lan"`
	WorkMode  C.WorkMode `yaml:"mode"`

	Proxy []map[string]any `yaml:"proxies"`
	Rules []string         `yaml:"rules"`
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
	g, rerr := ParseRawConfig(raw)
	if rerr != nil {
		return nil, rerr
	}
	// TODO check config
	return g, nil
}

func ParseRawConfig(raw *RawConfig) (*GlobalConfig, error) {
	g := &GlobalConfig{
		L: &LocalConfig{
			HttpProxyPort: C.DefaultPort,
		},
		Proxies:  map[string]Proxy{},
		WorkMode: raw.WorkMode,
	}
	if raw.HttpPort != 0 {
		g.L.HttpProxyPort = raw.HttpPort
		g.L.ProxyMode = C.HTTP
	}
	if raw.SocksPort != 0 {
		g.L.SocksProxyPort = raw.SocksPort
		g.L.ProxyMode = C.SOCKS
	}
	if raw.MixedPort != 0 {
		g.L.MixedProxyPort = raw.MixedPort
		g.L.ProxyMode = C.MIXED
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
	// parse rules
	r, err := parseRules(raw, g.Proxies)
	if err != nil {
		return nil, err
	}
	g.Rules = r
	return g, nil

}

func parseRules(cfg *RawConfig, proxies map[string]Proxy) ([]rules.Rule, error) {
	var rulesArray []rules.Rule
	rulesConfig := cfg.Rules

	// parse rulesArray
	for idx, line := range rulesConfig {
		rule := common.TrimArr(strings.Split(line, ","))
		var (
			payload string
			target  string
			params  []string
		)

		switch l := len(rule); {
		case l == 2:
			target = rule[1]
		case l == 3:
			payload = rule[1]
			target = rule[2]
		case l >= 4:
			payload = rule[1]
			target = rule[2]
			params = rule[3:]
		default:
			return nil, fmt.Errorf("rulesArray[%d] [%s] error: format invalid", idx, line)
		}

		if _, ok := proxies[target]; !ok {
			return nil, fmt.Errorf("rulesArray[%d] [%s] error: proxy [%s] not found", idx, line, target)
		}

		rule = common.TrimArr(rule)
		params = common.TrimArr(params)

		parsed, parseErr := rules.ParseRule(rule[0], payload, target, params)
		if parseErr != nil {
			return nil, fmt.Errorf("rulesArray[%d] [%s] error: %s", idx, line, parseErr.Error())
		}

		rulesArray = append(rulesArray, parsed)
	}

	return rulesArray, nil
}
