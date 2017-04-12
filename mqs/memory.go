package mqs

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// MemoryMQ message queue
type MemoryMQ struct {
	Queue        chan *SteelTask
	operateLinks map[string]*SteelTask
	//parallelChan  The maximum number of tasks that can be run at one time
	parallelChan chan int
	linksMutex   sync.RWMutex
}

// NewMemoryMQ New a MemoryMQ.
func NewMemoryMQ() *MemoryMQ {
	mq := MemoryMQ{
		Queue:        make(chan *SteelTask, *TasksCapability),
		operateLinks: make(map[string]*SteelTask, *TasksCapability),
		parallelChan: make(chan int, *MaxTasksParallel),
	}
	mq.start()
	return &mq
}

//start dispatch.
func (mq *MemoryMQ) start() {
	// start
	go func() {
		for {
			mq.parallelChan <- 1
			select {
			case task := <-mq.Queue:
				task.Lock()
				if task.State == TaskStateDelete {
					//if tag delete, just remove from link
					mq.linksMutex.Lock()
					delete(mq.operateLinks, task.Id)
					<-mq.parallelChan
					mq.linksMutex.Unlock()
				} else if task.State == TaskStatePendding {
					// Do the ture job
					ctx, cancel := context.WithCancel(context.Background())
					task.Cancel = cancel
					task.State = TaskStateRunning
					task.StartAt = time.Now()
					go func() {
						defer delete(mq.operateLinks, task.Id)
						defer cancel()
						defer func() {
							//Release paralle lock and delete form links(map).
							<-mq.parallelChan
							fmt.Printf("@@@@@ There are %d tasks left.\n", mq.Cnt())
						}()
						task.Func(ctx, task.Id, task.Args)
					}()
				} else {
					//FIXME error
					<-mq.parallelChan
					delete(mq.operateLinks, task.Id)
				}
				task.Unlock()
			}
		}
	}()
}

//Push Push a task into mq.
func (mq *MemoryMQ) Push(ctx context.Context, task *SteelTask) error {
	//mq.linksMutex.Lock()
	task.Lock()
	defer task.Unlock()
	task.State = TaskStatePendding
	mq.operateLinks[task.Id] = task
	//mq.linksMutex.Unlock()
	select {
	case <-ctx.Done():
		delete(mq.operateLinks, task.Id)
		return errors.New("push timeout")
	case mq.Queue <- task:
	}
	return nil
}

// Delete Delete the task from mq.
func (mq *MemoryMQ) Delete(taskid string) error {
	// mq.linksMutex.RLock()
	// defer mq.linksMutex.RUnlock()
	item, ok := mq.operateLinks[taskid]
	if !ok {
		return errors.New("task not found")
	}
	item.Lock()
	defer item.Unlock()
	// FIXME Just set the task state to 'delete' here,
	// but this task will not be delete from the queue.
	if item.State == TaskStatePendding {
		item.State = TaskStateDelete
	} else if item.State == TaskStateRunning {
		item.Cancel()
	}

	return nil
}

func (mq *MemoryMQ) Cnt() int {
	return len(mq.operateLinks)
}
