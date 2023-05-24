package common

import (
	"fmt"
	"strings"
)

func GenAddr(host string, port int) string {
	if host != "*" {
		return fmt.Sprintf("%s:%d", host, port)
	}
	return fmt.Sprintf("127.0.0.1:%d", port)
}

func TrimArr(arr []string) (r []string) {
	for _, e := range arr {
		r = append(r, strings.Trim(e, " "))
	}
	return
}
