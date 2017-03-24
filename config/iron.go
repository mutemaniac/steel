package config

import (
	"os"
)

const (
	ironFunciotnsServerEvn = "IRON_FUNCTIONS_SERVER_EVN"
)

var (
	IronFunciotnsServer string
)

func init() {
	IronFunciotnsServer = os.Getenv(ironFunciotnsServerEvn)
	if IronFunciotnsServer == "" {
		IronFunciotnsServer = `http://192.168.56.6:8080/v1`
	}
}
