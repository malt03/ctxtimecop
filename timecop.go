package timecop

import (
	"context"
	"time"
)

// Now wrapping time.Now()
func Now(ctx context.Context) time.Time {
	if isFreeze(ctx) {
		return freezeTime(ctx)
	}

	if isTravel(ctx) {
		return freezeTime(ctx).Add(time.Now().Sub(travelTime(ctx)))
	}

	return time.Now()
}

// Since wrapping time.Since()
func Since(ctx context.Context, t time.Time) time.Duration {
	return Now(ctx).Sub(t)
}

// WithFreeze timecop.Now() always return t with returned context.
func WithFreeze(ctx context.Context, t time.Time) context.Context {
	ctx = context.WithValue(ctx, freezeTimeKey{}, t)
	return context.WithValue(ctx, isFreezeKey{}, true)
}

// WithTravel timecop.Now() always return traveled time with returned context.
func WithTravel(ctx context.Context, t time.Time) context.Context {
	ctx = context.WithValue(ctx, freezeTimeKey{}, t)
	ctx = context.WithValue(ctx, travelTimeKey{}, time.Now())
	return context.WithValue(ctx, isTravelKey{}, true)
}

// WithReturn return freeze or travel
func WithReturn(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, isFreezeKey{}, false)
	return context.WithValue(ctx, isTravelKey{}, false)
}

type isFreezeKey struct{}
type freezeTimeKey struct{}
type isTravelKey struct{}
type travelTimeKey struct{}

func isFreeze(ctx context.Context) bool {
	return getBoolValue(ctx, isFreezeKey{})
}

func isTravel(ctx context.Context) bool {
	return getBoolValue(ctx, isTravelKey{})
}

func freezeTime(ctx context.Context) time.Time {
	return ctx.Value(freezeTimeKey{}).(time.Time)
}

func travelTime(ctx context.Context) time.Time {
	return ctx.Value(travelTimeKey{}).(time.Time)
}

func getBoolValue(ctx context.Context, key interface{}) bool {
	value := ctx.Value(key)
	if value == nil {
		return false
	}
	return value.(bool)
}
