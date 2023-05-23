package global

import "github.com/savion1024/wall/tunnel"

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

type Rules interface {
	RuleType() RuleType
	Match(c *tunnel.ConnContext) bool
	Adapter() string
	Payload() string
	ShouldResolveIP() bool
	ShouldFindProcess() bool
}
