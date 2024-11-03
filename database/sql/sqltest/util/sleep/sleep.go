// Package sleep provides a sleep function that respects a context.
package sleep

import (
	"context"
	"time"
)

// Sleep sleeps for a duration or until the context is done.
func Sleep(ctx context.Context, d time.Duration) error {
	timer := time.NewTimer(d)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return context.Cause(ctx)
	case <-timer.C:
		return nil
	}
}
