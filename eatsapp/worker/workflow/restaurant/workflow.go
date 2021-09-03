package restaurant

import (
	"time"

	"go.uber.org/zap"
	"trying/internal"

	"github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/activity/restaurant"
)

func init() {
	internal.RegisterWorkflow(OrderWorkflow)
}

// OrderWorkflow implements the restaurant order workflow.
func OrderWorkflow(ctx internal.Context, wfRunID string, orderID string, items []string) (time.Duration, error) {

	ao := internal.ActivityOptions{
		ScheduleToStartTimeout: time.Minute * 5,
		StartToCloseTimeout:    time.Minute * 15,
	}

	ctx = internal.WithActivityOptions(ctx, ao)
	err := internal.ExecuteActivity(ctx, restaurant.PlaceOrderActivity, wfRunID, orderID, items).Get(ctx, nil)
	if err != nil {
		internal.GetLogger(ctx).Error("Failed to send order to restaurant", zap.Error(err))
		return time.Minute * 0, err
	}

	var eta time.Duration
	err = internal.ExecuteActivity(ctx, restaurant.EstimateETAActivity, orderID).Get(ctx, &eta)
	if err != nil {
		internal.GetLogger(ctx).Error("Failed to estimate ETA for order ready", zap.Error(err))
		return time.Minute * 0, err
	}

	internal.GetLogger(ctx).Info("Completed PlaceOrder!")
	return eta, err
}
