package restaurant

import (
	"context"
	"time"
)

// func init() {
// 	workflow.RegisterActivity(EstimateETAActivity)
// }

// EstimateETAActivity implements the estimate eta activity.
func EstimateETAActivity(ctx context.Context, orderID string) (time.Duration, error) {
	return time.Minute, nil
}
