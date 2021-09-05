package eats

import (
	"fmt"
	"time"
	"net/http"
	"strings"
	"trying/helper"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"github.com/uber/cadence/common/types"
	//"github.com/uber/tchannel-go/crossdock/client"
	// "time"
)

// create creates a new eats order
func (h *EatsService) create(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	items := r.Form["item-id"]

	if len(items) == 0 {
		http.Error(w, "Order constains no items!", http.StatusUnprocessableEntity)
		return
	}

	execution, err := h.startOrderWorkflow(items)
	if err != nil {
		if strings.HasPrefix(err.Error(), "WorkflowExecutionAlreadyStartedError") {
			http.Redirect(w, r, "/eats-orders?error=order_exist", http.StatusFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := fmt.Sprintf("/eats-orders?id=%s&run_id=%s&page=eats-order-status", execution.WorkflowID, execution.RunID)
	http.Redirect(w, r, url, http.StatusFound)
}

// startOrderWorkflow starts the eats order workflow

func (h *EatsService) startOrderWorkflow(items []string) (*types.WorkflowExecution, error) {
	// THIS IS A PLACEHOLDER IMPLEMENTATION
	
	var wf helper.SampleHelper
	wf.SetupServiceConfig()
	startWorkflow(&wf)
	
	//return nil, fmt.Errorf("not implemented")
}

func startWorkflow(h *helper.SampleHelper) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "ubereats_" + uuid.New(),
		TaskList:                        "ApplicationName",
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}
	h.StartWorkflow(workflowOptions, "WorkflowName", "Cadence")
}
