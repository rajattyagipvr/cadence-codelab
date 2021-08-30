package main

import (
	"github.com/rajattyagipvr/cadence-codelab/common"
	"github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/activity/courier"
	"github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/activity/eats"
	"github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/activity/restaurant"
	"github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/workflow/courier"
	"github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/workflow/eats"
	"github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/workflow/restaurant"
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
