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
}
