package common

import "fmt"

func GenAddr(host string, port int) string {
	return fmt.Sprintf("127.0.0.1:%d", port)
}
