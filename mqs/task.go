package mqs

import (
	"context"
	"sync"
	"time"
)

const (
	//TaskStatePendding task state 'pendding'.
	TaskStatePendding = "pendding"
	//TaskStateRunning task state 'running'.
	TaskStateRunning = "running"
	//TaskStateStoped task state 'stoped'.
	TaskStateStoped = "stoped"
	//TaskStateDelete task state 'delete'.
	TaskStateDelete = "delete"
)

type callback func(ctx context.Context, Args interface{})
type cancel func()

//SteelTask  mq task
type SteelTask struct {
	Id       string
	State    string
	StartAt  time.Time
	Args     interface{}
	Callback callback
	Cancel   context.CancelFunc
	sync.Mutex
}
