package eats

import (
	"time"

	"go.uber.org/cadence/workflow"
)

func placeRestaurantOrder(ctx workflow.Context, orderID string, items []string) (time.Duration, error) {
	return time.Minute, nil
}
