package timecop_test

import (
	"context"
	"testing"
	"time"

	"github.com/malt03/ctxtimecop"
)

func TestFreeze(t *testing.T) {
	ctx0 := context.Background()
	ctx1 := context.Background()

	now := timecop.Now(ctx0)
	time.Sleep(50 * time.Millisecond)

	ctx0 = timecop.WithFreeze(ctx0, now)

	if timecop.Now(ctx0) != now {
		t.Errorf("Expected time is not %v.", now)
	}
	if timecop.Now(ctx1).Before(timecop.Now(ctx0).Add(49 * time.Millisecond)) {
		t.Errorf("timecop should not freeze another context.")
	}

	ctx0 = timecop.WithReturn(ctx0)

	if !timecop.Now(ctx0).Before(time.Now()) {
		t.Error("timecop should be reolve freezing.")
	}
}

func TestTravel(t *testing.T) {
	ctx := context.Background()
	now := timecop.Now(ctx)
	future := now.AddDate(1, 0, 0)
	ctx = timecop.WithTravel(ctx, future)

	if timecop.Now(ctx) == now {
		t.Errorf("Expected time is not %v.", now)
	}

	if !timecop.Now(ctx).After(future) {
		t.Errorf("Expected time should be greater than %v.", future)
	}
}
