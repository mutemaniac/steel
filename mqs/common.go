package mqs

import "flag"

// MaxTasksParallel The maximum number of tasks that can be run at one time
var MaxTasksParallel = flag.Int("tasks-parallel", 10, "The maximum number of tasks that can be run at one time")

// TasksCapability The maximum bumber of jobs that can be stored.
var TasksCapability = flag.Int("task-capability", 100, "The maximum bumber of jobs that can be stored")
