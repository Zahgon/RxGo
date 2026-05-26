package rxgo

import (
	"context"
	"sync"
)

type createIterable struct {
	next                   <-chan Item
	opts                   []Option
	subscribers            []chan Item
	mutex                  sync.RWMutex
	producerAlreadyCreated bool
}

func newCreateIterable(fs []Producer, opts ...Option) Iterable {
	_ = "STUB: not implemented"
	return *new(Iterable)
}

func (i *createIterable) Observe(opts ...Option) <-chan Item { _ = "STUB: not implemented"; return nil }

func (i *createIterable) connect(ctx context.Context) { _ = "STUB: not implemented"; return }

func (i *createIterable) produce(ctx context.Context) { _ = "STUB: not implemented"; return }
