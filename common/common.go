package common

import "fmt"

func GenAddr(host string, port int) string {
	if host != "*" {
		return fmt.Sprintf("%s:%d", host, port)
	}
	return fmt.Sprintf("127.0.0.1:%d", port)
}
