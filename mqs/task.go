package mqs

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
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

type Callback func(ctx context.Context, Args interface{})

//SteelTask  mq task
type SteelTask struct {
	Id       string
	State    string
	StartAt  time.Time
	Args     interface{}
	Callback Callback
	Cancel   context.CancelFunc
	sync.Mutex
}

// NewSteelTask generate a new steel task using arg & call back function.
func NewSteelTask(args interface{}, callback Callback) SteelTask {
	return SteelTask{
		Id:       uuid.New().String(),
		State:    TaskStatePendding,
		Args:     args,
		Callback: callback,
	}
}
