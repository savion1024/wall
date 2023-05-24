package rules

import (
	"github.com/savion1024/wall/tunnel"
	"strings"

	C "github.com/savion1024/wall/constant"
)

type Domain struct {
	domain  string
	adapter string
}

func (d *Domain) RuleType() C.RuleType {
	return C.Domain
}

func (d *Domain) Match(c *tunnel.ConnContext) bool {
	return c.RemoteConn.RemoteAddr().String() == d.domain
}

func (d *Domain) Adapter() string {
	return d.adapter
}

func (d *Domain) Payload() string {
	return d.domain
}

func (d *Domain) ShouldResolveIP() bool {
	return false
}

func (d *Domain) ShouldFindProcess() bool {
	return false
}

func NewDomain(domain string, adapter string) *Domain {
	return &Domain{
		domain:  strings.ToLower(domain),
		adapter: adapter,
	}
}
