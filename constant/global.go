package global

import (
	"encoding/json"
	"errors"
	"strings"
)

const (
	DefaultPort = 5678
)

type ProxyMode int
type WorkMode int

const (
	HTTP ProxyMode = iota
	SOCKS
	MIXED
)

const (
	Global WorkMode = iota
	Rule
	Direct
)

// ModeMapping is a mapping for Mode enum
var ModeMapping = map[string]WorkMode{
	Global.String(): Global,
	Rule.String():   Rule,
	Direct.String(): Direct,
}

func (p *ProxyMode) String() string {
	switch *p {
	case HTTP:
		return "HTTP"
	case SOCKS:
		return "SOCKS"
	case MIXED:
		return "MIXED"
	default:
		return "UNKNOWN"
	}
}

func (w WorkMode) String() string {
	switch w {
	case Global:
		return "global"
	case Rule:
		return "rule"
	case Direct:
		return "direct"
	default:
		return "Unknown"
	}
}

// MarshalJSON serialize Mode
func (w WorkMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(w.String())
}

// MarshalYAML serialize TunnelMode with yaml
func (w WorkMode) MarshalYAML() (any, error) {
	return w.String(), nil
}

// UnmarshalJSON unserialize Mode
func (w *WorkMode) UnmarshalJSON(data []byte) error {
	var tp string
	json.Unmarshal(data, &tp)
	mode, exist := ModeMapping[strings.ToLower(tp)]
	if !exist {
		return errors.New("invalid mode")
	}
	*w = mode
	return nil
}

// UnmarshalYAML unserialize Mode with yaml
func (w *WorkMode) UnmarshalYAML(unmarshal func(any) error) error {
	var tp string
	unmarshal(&tp)
	mode, exist := ModeMapping[strings.ToLower(tp)]
	if !exist {
		return errors.New("invalid mode")
	}
	*w = mode
	return nil
}

// Rule Type
const (
	Domain RuleType = iota
	DomainSuffix
	DomainKeyword
	GEOIP
	IPCIDR
	SrcIPCIDR
	SrcPort
	DstPort
	Process
	ProcessPath
	MATCH
)

type RuleType int

func (rt RuleType) String() string {
	switch rt {
	case Domain:
		return "Domain"
	case DomainSuffix:
		return "DomainSuffix"
	case DomainKeyword:
		return "DomainKeyword"
	case GEOIP:
		return "GeoIP"
	case IPCIDR:
		return "IPCIDR"
	case SrcIPCIDR:
		return "SrcIPCIDR"
	case SrcPort:
		return "SrcPort"
	case DstPort:
		return "DstPort"
	case Process:
		return "Process"
	case ProcessPath:
		return "ProcessPath"
	case MATCH:
		return "Match"
	default:
		return "Unknown"
	}
}
