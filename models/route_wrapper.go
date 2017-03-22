package models

import ironClient "github.com/iron-io/functions_go"

type Config map[string]string

type ExRouteWrapper struct {
	ironClient.Route
	AppName string `json:"appname" binding:"required"`
	Code    string `json:"code" binding:"required"`
	Runtime string `json:"runtime"  binding:"required"`
}
