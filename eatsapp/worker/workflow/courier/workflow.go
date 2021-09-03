package courier

import (
	"time"

	"github.com/rajattyagipvr/cadence-codelab/eatsapp/worker/activity/courier"
	"go.uber.org/zap"
	"trying/internal"
)

func init() {
	internal.RegisterWorkflow(OrderWorkflow)
}

// OrderWorkflow implements the deliver order workflow.
func OrderWorkflow(ctx internal.Context, orderID string) error {

	ao := internal.ActivityOptions{
		ScheduleToStartTimeout: time.Minute * 5,
		StartToCloseTimeout:    time.Minute * 15,
	}
	ctx = internal.WithActivityOptions(ctx, ao)

	for {
		err := internal.ExecuteActivity(ctx, courier.DispatchCourierActivity, orderID).Get(ctx, nil)
		if err != nil {
			// retry forever until a driver accepts the trip
			internal.GetLogger(ctx).Error("Failed to dispatch courier", zap.Error(err))
			continue
		}
		break
	}

	execution := internal.GetWorkflowInfo(ctx).WorkflowExecution
	err := internal.ExecuteActivity(ctx, courier.PickUpOrderActivity, execution, orderID).Get(ctx, nil)
	if err != nil {
		internal.GetLogger(ctx).Error("Failed to pick up order from restaurant", zap.Error(err))
		return err
	}

	err = waitForRestaurantPickupConfirmation(ctx, orderID)
	if err != nil {
		internal.GetLogger(ctx).Error("Failed to confirm pickup with restaurant", zap.Error(err))
		return err
	}

	err = internal.ExecuteActivity(ctx, courier.DeliverOrderActivity, orderID).Get(ctx, nil)
	if err != nil {
		internal.GetLogger(ctx).Error("Failed to complete delivery", zap.Error(err))
		return err
	}

	return nil
}
