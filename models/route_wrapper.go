package models

import "net/http"

type Config map[string]string

type RouteWrapper struct {
	AppName        string      `json:"app_name" binding:"required"`
	Path           string      `json:"path" binding:"required"`
	Code           string      `json:"code" binding:"required"`
	Language       string      `json:""  binding:"required"`
	Memory         uint64      `json:"memory"`
	Headers        http.Header `json:"headers"`
	Type           string      `json:"type"`
	Format         string      `json:"format"`
	MaxConcurrency int         `json:"max_concurrency"`
	Timeout        int32       `json:"timeout"`
	Config         `json:"config"`
}
