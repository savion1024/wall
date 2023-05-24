package rules

import (
	C "github.com/savion1024/wall/constant"
	"github.com/savion1024/wall/tunnel"
)

type Match struct {
	adapter string
}

func (f *Match) RuleType() C.RuleType {
	return C.MATCH
}

func (f *Match) Match(c *tunnel.ConnContext) bool {
	return true
}

func (f *Match) Adapter() string {
	return f.adapter
}

func (f *Match) Payload() string {
	return ""
}

func (f *Match) ShouldResolveIP() bool {
	return false
}

func (f *Match) ShouldFindProcess() bool {
	return false
}

func NewMatch(adapter string) *Match {
	return &Match{
		adapter: adapter,
	}
}
