package rxgo

import "context"

// Single is a observable with a single element.
type Single interface {
	Iterable
	Filter(apply Predicate, opts ...Option) OptionalSingle
	Get(opts ...Option) (Item, error)
	Map(apply Func, opts ...Option) Single
	Run(opts ...Option) Disposed
}

// SingleImpl implements Single.
type SingleImpl struct {
	parent   context.Context
	iterable Iterable
}

// Filter emits only those items from an Observable that pass a predicate test.
func (s *SingleImpl) Filter(apply Predicate, opts ...Option) OptionalSingle {
	_ = "STUB: not implemented"
	return *new(OptionalSingle)
}

// Get returns the item. The error returned is if the context has been cancelled.
// This method is blocking.
func (s *SingleImpl) Get(opts ...Option) (Item, error) {
	_ = "STUB: not implemented"
	return *new(Item), nil
}

// Map transforms the items emitted by a Single by applying a function to each item.
func (s *SingleImpl) Map(apply Func, opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type mapOperatorSingle struct {
	apply Func
}

func (op *mapOperatorSingle) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *mapOperatorSingle) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *mapOperatorSingle) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *mapOperatorSingle) gatherNext(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// Observe observes a Single by returning its channel.
func (s *SingleImpl) Observe(opts ...Option) <-chan Item { _ = "STUB: not implemented"; return nil }

type filterOperatorSingle struct {
	apply Predicate
}

func (op *filterOperatorSingle) next(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *filterOperatorSingle) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *filterOperatorSingle) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *filterOperatorSingle) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// Run creates an observer without consuming the emitted items.
	return
}

func (s *SingleImpl) Run(opts ...Option) Disposed { _ = "STUB: not implemented"; return *new(Disposed) }
