package rxgo

import (
	"context"
	"sync"
)

type channelIterable struct {
	next                   <-chan Item
	opts                   []Option
	subscribers            []chan Item
	mutex                  sync.RWMutex
	producerAlreadyCreated bool
}

func newChannelIterable(next <-chan Item, opts ...Option) Iterable {
	_ = "STUB: not implemented"
	return *new(Iterable)
}

func (i *channelIterable) Observe(opts ...Option) <-chan Item {
	_ = "STUB: not implemented"
	return nil
}

func (i *channelIterable) connect(ctx context.Context) { _ = "STUB: not implemented"; return }

func (i *channelIterable) produce(ctx context.Context) { _ = "STUB: not implemented"; return }
