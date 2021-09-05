package courier

import (
	"errors"

	"go.uber.org/cadence/workflow"
)

func waitForRestaurantPickupConfirmation(ctx workflow.Context, signalName string) error {
	return errors.New("not implemented")
}
