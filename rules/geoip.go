package rules

import (
	C "github.com/savion1024/wall/constant"
	"github.com/savion1024/wall/tunnel"
)

type GEOIP struct {
	country     string
	adapter     string
	noResolveIP bool
}

func (g *GEOIP) RuleType() C.RuleType {
	return C.GEOIP
}

func (g *GEOIP) Match(c *tunnel.ConnContext) bool {
	return true

}

func (g *GEOIP) Adapter() string {
	return g.adapter
}

func (g *GEOIP) Payload() string {
	return g.country
}

func (g *GEOIP) ShouldResolveIP() bool {
	return !g.noResolveIP
}

func (g *GEOIP) ShouldFindProcess() bool {
	return false
}

func NewGEOIP(country string, adapter string, noResolveIP bool) *GEOIP {
	geoip := &GEOIP{
		country:     country,
		adapter:     adapter,
		noResolveIP: noResolveIP,
	}

	return geoip
}
