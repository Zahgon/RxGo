package rxgo

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

// Infinite represents an infinite wait time
var Infinite int64 = -1

// Duration represents a duration
type Duration interface {
	duration() time.Duration
}

type duration struct {
	d time.Duration
}

func (d *duration) duration() time.Duration {
	_ = "STUB: not implemented"

	// WithDuration is a duration option
	return *new(time.Duration)
}

func WithDuration(d time.Duration) Duration { _ = "STUB: not implemented"; return *new(Duration) }

var tick = struct{}{}

type causalityDuration struct {
	fs []execution
}

type execution struct {
	f      func()
	isTick bool
}

func timeCausality(elems ...interface{}) (context.Context, Observable, Duration) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(Observable), *new(Duration)
}

func (d *causalityDuration) duration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type mockDuration struct {
	mock.Mock
}

func (m *mockDuration) duration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
