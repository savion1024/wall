package rules

import (
	"github.com/savion1024/wall/tunnel"
	"strconv"

	C "github.com/savion1024/wall/constant"
)

type Port struct {
	adapter  string
	port     string
	isSource bool
}

func (p *Port) RuleType() C.RuleType {
	if p.isSource {
		return C.SrcPort
	}
	return C.DstPort
}

func (p *Port) Match(c *tunnel.ConnContext) bool {
	//if p.isSource {
	//	return metadata.SrcPort == p.port
	//}
	//return metadata.DstPort == p.port
	return true
}

func (p *Port) Adapter() string {
	return p.adapter
}

func (p *Port) Payload() string {
	return p.port
}

func (p *Port) ShouldResolveIP() bool {
	return false
}

func (p *Port) ShouldFindProcess() bool {
	return false
}

func NewPort(port string, adapter string, isSource bool) (*Port, error) {
	_, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return nil, errPayload
	}
	return &Port{
		adapter:  adapter,
		port:     port,
		isSource: isSource,
	}, nil
}
