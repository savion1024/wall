package global

type ProxyMode int
type WorkMode int

const (
	HTTP ProxyMode = iota
	SOCKS
	MIXED
)

const (
	Global WorkMode = iota
	RULE
	DIRECT
)

const DefaultPort = "5678"
