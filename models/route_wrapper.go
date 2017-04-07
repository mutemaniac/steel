package models

import ironClient "github.com/mutemaniac/functions_go"

const (
	// TypeNone ...
	TypeNone = ""
	// TypeSync ...
	TypeSync = "sync"
	// TypeAsync ...
	TypeAsync = "async"
)
const (
	// FormatDefault ...
	FormatDefault = "default"
	// FormatHTTP ...
	FormatHTTP = "http"
)
const (
	defaultRouteTimeout = 30 // seconds
)

type Config map[string]string

type ExRouteWrapper struct {
	ironClient.Route `json:"route" binding:"required"`
	AppName          string `json:"appname" binding:"required"`
	Code             string `json:"code" binding:"required"`
	Runtime          string `json:"runtime"  binding:"required"`
}
type AsyncRouteWrapper struct {
	ExRouteWrapper
	Callback string `json:"callback"  binding:"required"`
}

// SetDefaults sets zeroed field to defaults.
func (r *ExRouteWrapper) SetDefaults(oldR *ironClient.Route) {
	if r.Route.Path == "" {
		r.Route.Path = oldR.Path
	}
	if r.Route.Image == "" {
		r.Route.Image = oldR.Image
	}

	if r.Route.Memory == 0 {
		r.Route.Memory = oldR.Memory
	}

	if r.Route.Type_ == TypeNone {
		r.Route.Type_ = oldR.Type_
	}

	if r.Route.Format == "" {
		r.Route.Format = oldR.Format
	}

	if r.Route.MaxConcurrency == 0 {
		r.Route.MaxConcurrency = oldR.MaxConcurrency
	}

	if r.Route.Headers == nil {
		r.Route.Headers = oldR.Headers
	}

	if r.Route.Config == nil {
		r.Route.Config = oldR.Config
	}

	if r.Route.Timeout == 0 {
		r.Route.Timeout = oldR.Timeout
	}
}
