package restaurant

import (
	"time"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"trying/worker/activity/restaurant"
)

// func init() {
// 	helper.RegisterWorkflow(OrderWorkflow)
// }

// OrderWorkflow implements the restaurant order workflow.
func OrderWorkflow(ctx workflow.Context, wfRunID string, orderID string, items []string) (time.Duration, error) {

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute * 5,
		StartToCloseTimeout:    time.Minute * 15,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)
	err := workflow.ExecuteActivity(ctx, restaurant.PlaceOrderActivity, wfRunID, orderID, items).Get(ctx, nil)
	if err != nil {
		workflow.GetLogger(ctx).Error("Failed to send order to restaurant", zap.Error(err))
		return time.Minute * 0, err
	}

	var eta time.Duration
	err = workflow.ExecuteActivity(ctx, restaurant.EstimateETAActivity, orderID).Get(ctx, &eta)
	if err != nil {
		workflow.GetLogger(ctx).Error("Failed to estimate ETA for order ready", zap.Error(err))
		return time.Minute * 0, err
	}

	workflow.GetLogger(ctx).Info("Completed PlaceOrder!")
	return eta, err
}
