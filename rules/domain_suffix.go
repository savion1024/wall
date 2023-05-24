package rules

import (
	"github.com/savion1024/wall/tunnel"
	"strings"

	C "github.com/savion1024/wall/constant"
)

type DomainSuffix struct {
	suffix  string
	adapter string
}

func (ds *DomainSuffix) RuleType() C.RuleType {
	return C.DomainSuffix
}

func (ds *DomainSuffix) Match(c *tunnel.ConnContext) bool {
	//domain := metadata.Host
	//return strings.HasSuffix(domain, "."+ds.suffix) || domain == ds.suffix
	return true
}

func (ds *DomainSuffix) Adapter() string {
	return ds.adapter
}

func (ds *DomainSuffix) Payload() string {
	return ds.suffix
}

func (ds *DomainSuffix) ShouldResolveIP() bool {
	return false
}

func (ds *DomainSuffix) ShouldFindProcess() bool {
	return false
}

func NewDomainSuffix(suffix string, adapter string) *DomainSuffix {
	return &DomainSuffix{
		suffix:  strings.ToLower(suffix),
		adapter: adapter,
	}
}
