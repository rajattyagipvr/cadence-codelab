package main

import (
	"time"
	"trying/helper"
	courieractivity "trying/worker/activity/courier"
	eatsactivity "trying/worker/activity/eats"
	restaurantactivity "trying/worker/activity/restaurant"
	courierworkflow "trying/worker/workflow/courier"
	eatsworkflow "trying/worker/workflow/eats"
	restaurantworkflow "trying/worker/workflow/restaurant"
	
	
	// "github.com/rajattyagipvr/cadence-codelab/common"
	// "github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/activity/courier"
	// "github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/activity/eats"
	// "github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/activity/restaurant"
	// "github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/workflow/courier"
	// "github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/workflow/eats"
	// "github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/workflow/restaurant"
	//"go.uber.org/cadence"

	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"


)

const (
	TaskListName = "cadence-bistro"
)

func main() {
	// runtime := common.NewRuntime()
	// // Configure worker options.
	// workerOptions :=helper.WorkerOptions{
	// 	MetricsScope: runtime.Scope,
	// 	Logger:       runtime.Logger,
	// }
	// runtime.StartWorkers(runtime.Config.DomainName, TaskListName, workerOptions)
	// select {}
	
	var h helper.SampleHelper
	h.SetupServiceConfig()
	registerWorkflowAndActivity(&h)
	startWorkers(&h)
	select {}


}


// This needs to be done as part of a bootstrap step when the process starts.
// The workers are supposed to be long running.
func startWorkers(h *helper.SampleHelper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.WorkerMetricScope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, "TestAppl", workerOptions)
}


func startWorkflow(h *helper.SampleHelper) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "helloworld_" + uuid.New(),
		TaskList:                        "TestAppl",
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}
	h.StartWorkflow(workflowOptions, eatsworkflow.OrderWorkflow, "EatsWorkflow")
}

func registerWorkflowAndActivity(
	h *helper.SampleHelper,
) {
	h.RegisterWorkflowWithAlias(courierworkflow.OrderWorkflow, "CourierWorkflow")
	h.RegisterActivity(courieractivity.PickUpOrderActivity)
	h.RegisterActivity(courieractivity.DispatchCourierActivity)
	h.RegisterActivity(courieractivity.DeliverOrderActivity)

	h.RegisterWorkflowWithAlias(restaurantworkflow.OrderWorkflow, "OrderWorkflow")
	h.RegisterActivity(restaurantactivity.PlaceOrderActivity)
	h.RegisterActivity(restaurantactivity.EstimateETAActivity)
	
	h.RegisterWorkflowWithAlias(eatsworkflow.OrderWorkflow, "EatsWorkflow")
	h.RegisterActivity(eatsactivity.ChargeOrderActivity)



}
