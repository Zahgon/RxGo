package rxgo

import (
	"context"
)

var emptyContext context.Context

// Option handles configurable options.
type Option interface {
	apply(*funcOption)
	toPropagate() bool
	isEagerObservation() bool
	getPool() (bool, int)
	buildChannel() chan Item
	buildContext(parent context.Context) context.Context
	getBackPressureStrategy() BackpressureStrategy
	getErrorStrategy() OnErrorStrategy
	isConnectable() bool
	isConnectOperation() bool
	isSerialized() (bool, func(interface{}) int)
}

type funcOption struct {
	f                    func(*funcOption)
	isBuffer             bool
	buffer               int
	ctx                  context.Context
	observation          ObservationStrategy
	pool                 int
	backPressureStrategy BackpressureStrategy
	onErrorStrategy      OnErrorStrategy
	propagate            bool
	connectable          bool
	connectOperation     bool
	serialized           func(interface{}) int
}

func (fdo *funcOption) toPropagate() bool { _ = "STUB: not implemented"; return false }

func (fdo *funcOption) isEagerObservation() bool { _ = "STUB: not implemented"; return false }

func (fdo *funcOption) getPool() (bool, int) { _ = "STUB: not implemented"; return false, 0 }

func (fdo *funcOption) buildChannel() chan Item { _ = "STUB: not implemented"; return nil }

func (fdo *funcOption) buildContext(parent context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (fdo *funcOption) getBackPressureStrategy() BackpressureStrategy {
	_ = "STUB: not implemented"
	return *new(BackpressureStrategy)
}

func (fdo *funcOption) getErrorStrategy() OnErrorStrategy {
	_ = "STUB: not implemented"
	return *new(OnErrorStrategy)
}

func (fdo *funcOption) isConnectable() bool { _ = "STUB: not implemented"; return false }

func (fdo *funcOption) isConnectOperation() bool { _ = "STUB: not implemented"; return false }

func (fdo *funcOption) apply(do *funcOption) { _ = "STUB: not implemented"; return }

func (fdo *funcOption) isSerialized() (bool, func(interface{}) int) {
	_ = "STUB: not implemented"
	return false, nil
}

func newFuncOption(f func(*funcOption)) *funcOption { _ = "STUB: not implemented"; return nil }

func parseOptions(opts ...Option) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBufferedChannel allows to configure the capacity of a buffered channel.
func WithBufferedChannel(capacity int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithContext allows to pass a context.
func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithObservationStrategy uses the eager observation mode meaning consuming the items even without subscription.
func WithObservationStrategy(strategy ObservationStrategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPool allows to specify an execution pool.
func WithPool(pool int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithCPUPool allows to specify an execution pool based on the number of logical CPUs.
func WithCPUPool() Option { _ = "STUB: not implemented"; return *new(Option) }

// WithBackPressureStrategy sets the back pressure strategy: drop or block.
func WithBackPressureStrategy(strategy BackpressureStrategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithErrorStrategy defines how an observable should deal with error.
// This strategy is propagated to the parent observable.
func WithErrorStrategy(strategy OnErrorStrategy) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithPublishStrategy converts an ordinary Observable into a connectable Observable.
func WithPublishStrategy() Option { _ = "STUB: not implemented"; return *new(Option) }

// Serialize forces an Observable to make serialized calls and to be well-behaved.
func Serialize(identifier func(interface{}) int) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func connect() Option { _ = "STUB: not implemented"; return *new(Option) }
