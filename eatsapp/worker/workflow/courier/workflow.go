package courier

import (
	"time"
	"trying/worker/activity/courier"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// func init() {
// 	.RegisterWorkflow(OrderWorkflow)
// }

// OrderWorkflow implements the deliver order workflow.
func OrderWorkflow(ctx workflow.Context, orderID string) error {

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute * 5,
		StartToCloseTimeout:    time.Minute * 15,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	for {
		err := workflow.ExecuteActivity(ctx, courier.DispatchCourierActivity, orderID).Get(ctx, nil)
		if err != nil {
			// retry forever until a driver accepts the trip
			workflow.GetLogger(ctx).Error("Failed to dispatch courier", zap.Error(err))
			continue
		}
		break
	}

	//execution := workflow.Execution
	err := workflow.ExecuteActivity(ctx, courier.PickUpOrderActivity, orderID).Get(ctx, nil)
	if err != nil {
		workflow.GetLogger(ctx).Error("Failed to pick up order from restaurant", zap.Error(err))
		return err
	}

	err = waitForRestaurantPickupConfirmation(ctx, orderID)
	if err != nil {
		workflow.GetLogger(ctx).Error("Failed to confirm pickup with restaurant", zap.Error(err))
		return err
	}

	err = workflow.ExecuteActivity(ctx, courier.DeliverOrderActivity, orderID).Get(ctx, nil)
	if err != nil {
		workflow.GetLogger(ctx).Error("Failed to complete delivery", zap.Error(err))
		return err
	}

	return nil
}
