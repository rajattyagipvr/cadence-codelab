package eats

import (
	"context"
	"time"

	"go.uber.org/zap"
	"trying/internal"
)

func init() {
	internal.RegisterActivity(ChargeOrderActivity)
}

// ChargeOrderActivity implements the change order activity.
func ChargeOrderActivity(ctx context.Context, orderID string) error {
	time.Sleep(time.Second * 5)
	internal.GetActivityLogger(ctx).Info("Charged customer for order!", zap.String("order", orderID))
	return nil
}
