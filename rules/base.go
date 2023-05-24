package rules

import (
	"errors"
	C "github.com/savion1024/wall/constant"
	"github.com/savion1024/wall/tunnel"
)

var (
	errPayload = errors.New("payload error")

	noResolve = "no-resolve"
)

func HasNoResolve(params []string) bool {
	for _, p := range params {
		if p == noResolve {
			return true
		}
	}
	return false
}

type Rule interface {
	RuleType() C.RuleType
	Match(c *tunnel.ConnContext) bool
	Adapter() string
	Payload() string
	ShouldResolveIP() bool
	ShouldFindProcess() bool
}
