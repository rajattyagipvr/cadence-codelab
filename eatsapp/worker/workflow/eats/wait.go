package eats

import (
	"errors"
	"time"

	"go.uber.org/cadence/workflow"
)

func waitForRestaurant(ctx workflow.Context, signalName string, eta time.Duration) error {
	return errors.New("not implemented")
}
