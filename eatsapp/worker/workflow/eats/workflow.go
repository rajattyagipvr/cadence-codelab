package eats

import (
	"go.uber.org/zap"
	"trying/internal"
)

func init() {
	internal.RegisterWorkflow(OrderWorkflow)
}

// OrderWorkflow implements the eats order workflow.
func OrderWorkflow(ctx internal.Context, orderID string, items []string) error {

	internal.GetLogger(ctx).Info("Received order", zap.Strings("items", items))

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

	internal.GetLogger(ctx).Info("Completed order", zap.String("order", orderID))
	return nil
}
