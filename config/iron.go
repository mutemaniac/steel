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
		IronFunciotnsServer = `http://52.80.17.251:8083`
	}
}
