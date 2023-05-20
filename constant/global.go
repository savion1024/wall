package constant

type WorkMode int

const (
	HTTP WorkMode = iota
	SOCKS
	MIXED
)

const DefaultPort = 5678
