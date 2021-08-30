package main

import (
	"github.com/venkat1109/cadence-codelab/common"
	"github.com/venkat1109/cadence-codelab/eatsapp/worker/activity/courier"
	"github.com/venkat1109/cadence-codelab/eatsapp/worker/activity/eats"
	"github.com/venkat1109/cadence-codelab/eatsapp/worker/activity/restaurant"
	"github.com/venkat1109/cadence-codelab/eatsapp/worker/workflow/courier"
	"github.com/venkat1109/cadence-codelab/eatsapp/worker/workflow/eats"
	"github.com/venkat1109/cadence-codelab/eatsapp/worker/workflow/restaurant"
	"go.uber.org/cadence"
)

const (
	TaskListName = "cadence-bistro"
)

func main() {
	runtime := common.NewRuntime()
	// Configure worker options.
	workerOptions := cadence.WorkerOptions{
		MetricsScope: runtime.Scope,
		Logger:       runtime.Logger,
	}
	runtime.StartWorkers(runtime.Config.DomainName, TaskListName, workerOptions)
	select {}
}
