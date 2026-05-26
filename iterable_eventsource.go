package rxgo

import (
	"context"
	"sync"
)

type eventSourceIterable struct {
	sync.RWMutex
	observers []chan Item
	disposed  bool
	opts      []Option
}

func newEventSourceIterable(ctx context.Context, next <-chan Item, strategy BackpressureStrategy, opts ...Option) Iterable {
	_ = "STUB: not implemented"
	return *new(Iterable)
}

func (i *eventSourceIterable) closeAllObservers() { _ = "STUB: not implemented"; return }

func (i *eventSourceIterable) Observe(opts ...Option) <-chan Item {
	_ = "STUB: not implemented"
	return nil
}
