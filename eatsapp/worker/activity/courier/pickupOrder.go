package courier

import (
	"context"
	"errors"

	"trying/internal"
)

func init() {
	internal.RegisterActivity(PickUpOrderActivity)
}

// PickUpOrderActivity implements the pick-up order activity.
func PickUpOrderActivity(ctx context.Context, execution internal.WorkflowExecution, orderID string) (string, error) {
	return "", errors.New("not implemented")
}

func notifyRestaurant(execution internal.WorkflowExecution, orderID string) error {
	url := "http://localhost:8090/restaurant?action=p_sig&id=" + orderID +
		"&workflow_id=" + execution.ID + "&run_id=" + execution.RunID
	return sendPatch(url)
}

func pickup(orderID string, taskToken string) error {
	url := "http://localhost:8090/courier?action=p_token&id=" + orderID + "&task_token=" + taskToken
	return sendPatch(url)
}
