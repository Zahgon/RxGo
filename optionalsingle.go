package rxgo

import "context"

// OptionalSingleEmpty is the constant returned when an OptionalSingle is empty.
var OptionalSingleEmpty = Item{}

// OptionalSingle is an optional single.
type OptionalSingle interface {
	Iterable
	Get(opts ...Option) (Item, error)
	Map(apply Func, opts ...Option) OptionalSingle
	Run(opts ...Option) Disposed
}

// OptionalSingleImpl implements OptionalSingle.
type OptionalSingleImpl struct {
	parent   context.Context
	iterable Iterable
}

// Get returns the item or rxgo.OptionalEmpty. The error returned is if the context has been cancelled.
// This method is blocking.
func (o *OptionalSingleImpl) Get(opts ...Option) (Item, error) {
	_ = "STUB: not implemented"
	return *new(Item), nil
}

// Map transforms the items emitted by an OptionalSingle by applying a function to each item.
func (o *OptionalSingleImpl) Map(apply Func, opts ...Option) OptionalSingle {
	_ = "STUB: not implemented"
	return *new(OptionalSingle)
}

// Observe observes an OptionalSingle by returning its channel.
func (o *OptionalSingleImpl) Observe(opts ...Option) <-chan Item {
	_ = "STUB: not implemented"
	return nil
}

type mapOperatorOptionalSingle struct {
	apply Func
}

func (op *mapOperatorOptionalSingle) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *mapOperatorOptionalSingle) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *mapOperatorOptionalSingle) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *mapOperatorOptionalSingle) gatherNext(_ context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// Run creates an observer without consuming the emitted items.
func (o *OptionalSingleImpl) Run(opts ...Option) Disposed {
	_ = "STUB: not implemented"
	return *new(Disposed)
}
