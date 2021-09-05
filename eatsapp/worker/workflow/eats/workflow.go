package eats

import (
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// func init() {
// 	helper.RegisterWorkflow(OrderWorkflow)
// }

// OrderWorkflow implements the eats order workflow.
func OrderWorkflow(ctx workflow.Context, orderID string, items []string) error {

	workflow.GetLogger(ctx).Info("Received order", zap.Strings("items", items))

	restaurantEta, err := placeRestaurantOrder(ctx, orderID, items)
	if err != nil {
		return err
	}

	err = waitForRestaurant(ctx, orderID, restaurantEta)
	if err != nil {
		return err
	}

	err = deliverOrder(ctx, orderID)
	if err != nil {
		return err
	}

	err = chargeOrder(ctx, orderID)
	if err != nil {
		return err
	}

	workflow.GetLogger(ctx).Info("Completed order", zap.String("order", orderID))
	return nil
}
