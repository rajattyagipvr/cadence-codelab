package eats

import (
	"context"
	"time"

	"go.uber.org/cadence/activity"
	"go.uber.org/zap"
)

func init() {
	activity.RegisterActivity(ChargeOrderActivity)
}

// ChargeOrderActivity implements the change order activity.
func ChargeOrderActivity(ctx context.Context, orderID string) error {
	time.Sleep(time.Second * 5)
	cadence.GetActivityLogger(ctx).Info("Charged customer for order!", zap.String("order", orderID))
	return nil
}
