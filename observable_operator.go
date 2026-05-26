package rxgo

import (
	"container/ring"
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
)

// All determines whether all items emitted by an Observable meet some criteria.
func (o *ObservableImpl) All(predicate Predicate, opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type allOperator struct {
	predicate Predicate
	all       bool
}

func (op *allOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *allOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *allOperator) end(ctx context.Context, dst chan<- Item) { _ = "STUB: not implemented"; return }

func (op *allOperator) gatherNext(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// AverageFloat32 calculates the average of numbers emitted by an Observable and emits the average float32.
func (o *ObservableImpl) AverageFloat32(opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type averageFloat32Operator struct {
	sum   float32
	count float32
}

func (op *averageFloat32Operator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageFloat32Operator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageFloat32Operator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *averageFloat32Operator) gatherNext(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// AverageFloat64 calculates the average of numbers emitted by an Observable and emits the average float64.
func (o *ObservableImpl) AverageFloat64(opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type averageFloat64Operator struct {
	sum   float64
	count float64
}

func (op *averageFloat64Operator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageFloat64Operator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageFloat64Operator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *averageFloat64Operator) gatherNext(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// AverageInt calculates the average of numbers emitted by an Observable and emits the average int.
func (o *ObservableImpl) AverageInt(opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type averageIntOperator struct {
	sum   int
	count int
}

func (op *averageIntOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageIntOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageIntOperator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *averageIntOperator) gatherNext(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// AverageInt8 calculates the average of numbers emitted by an Observable and emits the≤ average int8.
func (o *ObservableImpl) AverageInt8(opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type averageInt8Operator struct {
	sum   int8
	count int8
}

func (op *averageInt8Operator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageInt8Operator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageInt8Operator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *averageInt8Operator) gatherNext(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// AverageInt16 calculates the average of numbers emitted by an Observable and emits the average int16.
func (o *ObservableImpl) AverageInt16(opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type averageInt16Operator struct {
	sum   int16
	count int16
}

func (op *averageInt16Operator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageInt16Operator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageInt16Operator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *averageInt16Operator) gatherNext(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// AverageInt32 calculates the average of numbers emitted by an Observable and emits the average int32.
func (o *ObservableImpl) AverageInt32(opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type averageInt32Operator struct {
	sum   int32
	count int32
}

func (op *averageInt32Operator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageInt32Operator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageInt32Operator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *averageInt32Operator) gatherNext(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// AverageInt64 calculates the average of numbers emitted by an Observable and emits this average int64.
func (o *ObservableImpl) AverageInt64(opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type averageInt64Operator struct {
	sum   int64
	count int64
}

func (op *averageInt64Operator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageInt64Operator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *averageInt64Operator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *averageInt64Operator) gatherNext(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// BackOffRetry implements a backoff retry if a source Observable sends an error, resubscribe to it in the hopes that it will complete without error.
// Cannot be run in parallel.
func (o *ObservableImpl) BackOffRetry(backOffCfg backoff.BackOff, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// BufferWithCount returns an Observable that emits buffers of items it collects
// from the source Observable.
// The resulting Observable emits buffers every skip items, each containing a slice of count items.
// When the source Observable completes or encounters an error,
// the resulting Observable emits the current buffer and propagates
// the notification from the source Observable.
func (o *ObservableImpl) BufferWithCount(count int, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type bufferWithCountOperator struct {
	count  int
	iCount int
	buffer []interface{}
}

func (op *bufferWithCountOperator) next(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *bufferWithCountOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *bufferWithCountOperator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *bufferWithCountOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// BufferWithTime returns an Observable that emits buffers of items it collects from the source
	// Observable. The resulting Observable starts a new buffer periodically, as determined by the
	// timeshift argument. It emits each buffer after a fixed timespan, specified by the timespan argument.
	// When the source Observable completes or encounters an error, the resulting Observable emits
	// the current buffer and propagates the notification from the source Observable.
	return
}

func (o *ObservableImpl) BufferWithTime(timespan Duration, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// BufferWithTimeOrCount returns an Observable that emits buffers of items it collects from the source
// Observable either from a given count or at a given time interval.
func (o *ObservableImpl) BufferWithTimeOrCount(timespan Duration, count int, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Connect instructs a connectable Observable to begin emitting items to its subscribers.
func (o *ObservableImpl) Connect(ctx context.Context) (context.Context, Disposable) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(Disposable)
}

// Contains determines whether an Observable emits a particular item or not.
func (o *ObservableImpl) Contains(equal Predicate, opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type containsOperator struct {
	equal    Predicate
	contains bool
}

func (op *containsOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *containsOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *containsOperator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *containsOperator) gatherNext(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// Count counts the number of items emitted by the source Observable and emit only this value.
func (o *ObservableImpl) Count(opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type countOperator struct {
	count int64
}

func (op *countOperator) next(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *countOperator) err(_ context.Context, _ Item, _ chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *countOperator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *countOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// Debounce only emits an item from an Observable if a particular timespan has passed without it emitting another item.
	return
}

func (o *ObservableImpl) Debounce(timespan Duration, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// DefaultIfEmpty returns an Observable that emits the items emitted by the source
// Observable or a specified default item if the source Observable is empty.
func (o *ObservableImpl) DefaultIfEmpty(defaultValue interface{}, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type defaultIfEmptyOperator struct {
	defaultValue interface{}
	empty        bool
}

func (op *defaultIfEmptyOperator) next(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *defaultIfEmptyOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *defaultIfEmptyOperator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *defaultIfEmptyOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// Distinct suppresses duplicate items in the original Observable and returns
	// a new Observable.
	return
}

func (o *ObservableImpl) Distinct(apply Func, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type distinctOperator struct {
	apply  Func
	keyset map[interface{}]interface{}
}

func (op *distinctOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *distinctOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *distinctOperator) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *distinctOperator) gatherNext(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// DistinctUntilChanged suppresses consecutive duplicate items in the original Observable.
// Cannot be run in parallel.
func (o *ObservableImpl) DistinctUntilChanged(apply Func, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type distinctUntilChangedOperator struct {
	apply   Func
	current interface{}
}

func (op *distinctUntilChangedOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *distinctUntilChangedOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *distinctUntilChangedOperator) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *distinctUntilChangedOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// DoOnCompleted registers a callback action that will be called once the Observable terminates.
	return
}

func (o *ObservableImpl) DoOnCompleted(completedFunc CompletedFunc, opts ...Option) Disposed {
	_ = "STUB: not implemented"
	return *new(Disposed)
}

// DoOnError registers a callback action that will be called if the Observable terminates abnormally.
func (o *ObservableImpl) DoOnError(errFunc ErrFunc, opts ...Option) Disposed {
	_ = "STUB: not implemented"
	return *new(Disposed)
}

// DoOnNext registers a callback action that will be called on each item emitted by the Observable.
func (o *ObservableImpl) DoOnNext(nextFunc NextFunc, opts ...Option) Disposed {
	_ = "STUB: not implemented"
	return *new(Disposed)
}

// ElementAt emits only item n emitted by an Observable.
// Cannot be run in parallel.
func (o *ObservableImpl) ElementAt(index uint, opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type elementAtOperator struct {
	index     uint
	takeCount int
	sent      bool
}

func (op *elementAtOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *elementAtOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *elementAtOperator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *elementAtOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// Error returns the eventual Observable error.
	// This method is blocking.
	return
}

func (o *ObservableImpl) Error(opts ...Option) error { _ = "STUB: not implemented"; return nil }

// Errors returns an eventual list of Observable errors.
// This method is blocking
func (o *ObservableImpl) Errors(opts ...Option) []error { _ = "STUB: not implemented"; return nil }

// Filter emits only those items from an Observable that pass a predicate test.
func (o *ObservableImpl) Filter(apply Predicate, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type filterOperator struct {
	apply Predicate
}

func (op *filterOperator) next(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *filterOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *filterOperator) end(_ context.Context, _ chan<- Item) { _ = "STUB: not implemented"; return }

func (op *filterOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// Find emits the first item passing a predicate then complete.
	return
}

func (o *ObservableImpl) Find(find Predicate, opts ...Option) OptionalSingle {
	_ = "STUB: not implemented"
	return *new(OptionalSingle)
}

type findOperator struct {
	find Predicate
}

func (op *findOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *findOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *findOperator) end(_ context.Context, _ chan<- Item) { _ = "STUB: not implemented"; return }

func (op *findOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// First returns new Observable which emit only first item.
	// Cannot be run in parallel.
	return
}

func (o *ObservableImpl) First(opts ...Option) OptionalSingle {
	_ = "STUB: not implemented"
	return *new(OptionalSingle)
}

type firstOperator struct{}

func (op *firstOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *firstOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *firstOperator) end(_ context.Context, _ chan<- Item) { _ = "STUB: not implemented"; return }

func (op *firstOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// FirstOrDefault returns new Observable which emit only first item.
	// If the observable fails to emit any items, it emits a default value.
	// Cannot be run in parallel.
	return
}

func (o *ObservableImpl) FirstOrDefault(defaultValue interface{}, opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type firstOrDefaultOperator struct {
	defaultValue interface{}
	sent         bool
}

func (op *firstOrDefaultOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *firstOrDefaultOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *firstOrDefaultOperator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *firstOrDefaultOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// FlatMap transforms the items emitted by an Observable into Observables, then flatten the emissions from those into a single Observable.
	return
}

func (o *ObservableImpl) FlatMap(apply ItemToObservable, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// ForEach subscribes to the Observable and receives notifications for each element.
func (o *ObservableImpl) ForEach(nextFunc NextFunc, errFunc ErrFunc, completedFunc CompletedFunc, opts ...Option) Disposed {
	_ = "STUB: not implemented"
	return *new(Disposed)
}

// IgnoreElements ignores all items emitted by the source ObservableSource except for the errors.
// Cannot be run in parallel.
func (o *ObservableImpl) IgnoreElements(opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type ignoreElementsOperator struct{}

func (op *ignoreElementsOperator) next(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *ignoreElementsOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *ignoreElementsOperator) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *ignoreElementsOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// Returns absolute value for int64
	return
}

func abs(n int64) int64 { _ = "STUB: not implemented"; return 0 }

// Join combines items emitted by two Observables whenever an item from one Observable is emitted during
// a time window defined according to an item emitted by the other Observable.
// The time is extracted using a timeExtractor function.
func (o *ObservableImpl) Join(joiner Func2, right Observable, timeExtractor func(interface{}) time.Time, window Duration, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// GroupBy divides an Observable into a set of Observables that each emit a different group of items from the original Observable, organized by key.
func (o *ObservableImpl) GroupBy(length int, distribution func(Item) int, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// GroupedObservable is the observable type emitted by the GroupByDynamic operator.
type GroupedObservable struct {
	Observable
	// Key is the distribution key
	Key string
}

// GroupByDynamic divides an Observable into a dynamic set of Observables that each emit GroupedObservable from the original Observable, organized by key.
func (o *ObservableImpl) GroupByDynamic(distribution func(Item) string, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Last returns a new Observable which emit only last item.
// Cannot be run in parallel.
func (o *ObservableImpl) Last(opts ...Option) OptionalSingle {
	_ = "STUB: not implemented"
	return *new(OptionalSingle)
}

type lastOperator struct {
	last  Item
	empty bool
}

func (op *lastOperator) next(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *lastOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *lastOperator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *lastOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// LastOrDefault returns a new Observable which emit only last item.
	// If the observable fails to emit any items, it emits a default value.
	// Cannot be run in parallel.
	return
}

func (o *ObservableImpl) LastOrDefault(defaultValue interface{}, opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type lastOrDefaultOperator struct {
	defaultValue interface{}
	last         Item
	empty        bool
}

func (op *lastOrDefaultOperator) next(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *lastOrDefaultOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *lastOrDefaultOperator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *lastOrDefaultOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// Map transforms the items emitted by an Observable by applying a function to each item.
	return
}

func (o *ObservableImpl) Map(apply Func, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type mapOperator struct {
	apply Func
}

func (op *mapOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *mapOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *mapOperator) end(_ context.Context, _ chan<- Item) { _ = "STUB: not implemented"; return }

func (op *mapOperator) gatherNext(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// Marshal transforms the items emitted by an Observable by applying a marshalling to each item.
func (o *ObservableImpl) Marshal(marshaller Marshaller, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Max determines and emits the maximum-valued item emitted by an Observable according to a comparator.
func (o *ObservableImpl) Max(comparator Comparator, opts ...Option) OptionalSingle {
	_ = "STUB: not implemented"
	return *new(OptionalSingle)
}

type maxOperator struct {
	comparator Comparator
	empty      bool
	max        interface{}
}

func (op *maxOperator) next(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *maxOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *maxOperator) end(ctx context.Context, dst chan<- Item) { _ = "STUB: not implemented"; return }

func (op *maxOperator) gatherNext(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// Min determines and emits the minimum-valued item emitted by an Observable according to a comparator.
func (o *ObservableImpl) Min(comparator Comparator, opts ...Option) OptionalSingle {
	_ = "STUB: not implemented"
	return *new(OptionalSingle)
}

type minOperator struct {
	comparator Comparator
	empty      bool
	max        interface{}
}

func (op *minOperator) next(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *minOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *minOperator) end(ctx context.Context, dst chan<- Item) { _ = "STUB: not implemented"; return }

func (op *minOperator) gatherNext(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// Observe observes an Observable by returning its channel.
func (o *ObservableImpl) Observe(opts ...Option) <-chan Item { _ = "STUB: not implemented"; return nil }

// OnErrorResumeNext instructs an Observable to pass control to another Observable rather than invoking
// onError if it encounters an error.
func (o *ObservableImpl) OnErrorResumeNext(resumeSequence ErrorToObservable, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type onErrorResumeNextOperator struct {
	resumeSequence ErrorToObservable
}

func (op *onErrorResumeNextOperator) next(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *onErrorResumeNextOperator) err(_ context.Context, item Item, _ chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *onErrorResumeNextOperator) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *onErrorResumeNextOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// OnErrorReturn instructs an Observable to emit an item (returned by a specified function)
	// rather than invoking onError if it encounters an error.
	return
}

func (o *ObservableImpl) OnErrorReturn(resumeFunc ErrorFunc, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type onErrorReturnOperator struct {
	resumeFunc ErrorFunc
}

func (op *onErrorReturnOperator) next(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *onErrorReturnOperator) err(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *onErrorReturnOperator) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *onErrorReturnOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// OnErrorReturnItem instructs on Observable to emit an item if it encounters an error.
	return
}

func (o *ObservableImpl) OnErrorReturnItem(resume interface{}, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type onErrorReturnItemOperator struct {
	resume interface{}
}

func (op *onErrorReturnItemOperator) next(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *onErrorReturnItemOperator) err(ctx context.Context, _ Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *onErrorReturnItemOperator) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *onErrorReturnItemOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// Reduce applies a function to each item emitted by an Observable, sequentially, and emit the final value.
	return
}

func (o *ObservableImpl) Reduce(apply Func2, opts ...Option) OptionalSingle {
	_ = "STUB: not implemented"
	return *new(OptionalSingle)
}

type reduceOperator struct {
	apply Func2
	acc   interface{}
	empty bool
}

func (op *reduceOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *reduceOperator) err(_ context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *reduceOperator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *reduceOperator) gatherNext(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

// Repeat returns an Observable that repeats the sequence of items emitted by the source Observable
// at most count times, at a particular frequency.
// Cannot run in parallel.
func (o *ObservableImpl) Repeat(count int64, frequency Duration, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type repeatOperator struct {
	count     int64
	frequency Duration
	seq       []Item
}

func (op *repeatOperator) next(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *repeatOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *repeatOperator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *repeatOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// Retry retries if a source Observable sends an error, resubscribe to it in the hopes that it will complete without error.
	// Cannot be run in parallel.
	return
}

func (o *ObservableImpl) Retry(count int, shouldRetry func(error) bool, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Run creates an Observer without consuming the emitted items.
func (o *ObservableImpl) Run(opts ...Option) Disposed {
	_ = "STUB: not implemented"
	return *new(Disposed)
}

// Sample returns an Observable that emits the most recent items emitted by the source
// Iterable whenever the input Iterable emits an item.
func (o *ObservableImpl) Sample(iterable Iterable, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Scan apply a Func2 to each item emitted by an Observable, sequentially, and emit each successive value.
// Cannot be run in parallel.
func (o *ObservableImpl) Scan(apply Func2, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type scanOperator struct {
	apply   Func2
	current interface{}
}

func (op *scanOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *scanOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *scanOperator) end(_ context.Context, _ chan<- Item) { _ = "STUB: not implemented"; return }

func (op *scanOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// Compares first items of two sequences and returns true if they are equal and false if
	// they are not. Besides, it returns two new sequences - input sequences without compared items.
	return
}

func popAndCompareFirstItems(
	inputSequence1 []interface{},
	inputSequence2 []interface{}) (bool, []interface{}, []interface{}) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Send sends the items to a given channel.
func (o *ObservableImpl) Send(output chan<- Item, opts ...Option) {
	_ = "STUB: not implemented"
	return
}

// SequenceEqual emits true if an Observable and the input Observable emit the same items,
// in the same order, with the same termination state. Otherwise, it emits false.
func (o *ObservableImpl) SequenceEqual(iterable Iterable, opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

// Serialize forces an Observable to make serialized calls and to be well-behaved.
func (o *ObservableImpl) Serialize(from int, identifier func(interface{}) int, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Skip suppresses the first n items in the original Observable and
// returns a new Observable with the rest items.
// Cannot be run in parallel.
func (o *ObservableImpl) Skip(nth uint, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type skipOperator struct {
	nth       uint
	skipCount int
}

func (op *skipOperator) next(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *skipOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *skipOperator) end(_ context.Context, _ chan<- Item) { _ = "STUB: not implemented"; return }

func (op *skipOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// SkipLast suppresses the last n items in the original Observable and
	// returns a new Observable with the rest items.
	// Cannot be run in parallel.
	return
}

func (o *ObservableImpl) SkipLast(nth uint, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type skipLastOperator struct {
	nth       uint
	skipCount int
}

func (op *skipLastOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *skipLastOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *skipLastOperator) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *skipLastOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// SkipWhile discard items emitted by an Observable until a specified condition becomes false.
	// Cannot be run in parallel.
	return
}

func (o *ObservableImpl) SkipWhile(apply Predicate, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type skipWhileOperator struct {
	apply Predicate
	skip  bool
}

func (op *skipWhileOperator) next(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *skipWhileOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *skipWhileOperator) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *skipWhileOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// StartWith emits a specified Iterable before beginning to emit the items from the source Observable.
	return
}

func (o *ObservableImpl) StartWith(iterable Iterable, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// SumFloat32 calculates the average of float32 emitted by an Observable and emits a float32.
func (o *ObservableImpl) SumFloat32(opts ...Option) OptionalSingle {
	_ = "STUB: not implemented"
	return *new(OptionalSingle)
}

// SumFloat64 calculates the average of float64 emitted by an Observable and emits a float64.
func (o *ObservableImpl) SumFloat64(opts ...Option) OptionalSingle {
	_ = "STUB: not implemented"
	return *new(OptionalSingle)
}

// SumInt64 calculates the average of integers emitted by an Observable and emits an int64.
func (o *ObservableImpl) SumInt64(opts ...Option) OptionalSingle {
	_ = "STUB: not implemented"
	return *new(OptionalSingle)
}

// Take emits only the first n items emitted by an Observable.
// Cannot be run in parallel.
func (o *ObservableImpl) Take(nth uint, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type takeOperator struct {
	nth       uint
	takeCount int
}

func (op *takeOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *takeOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *takeOperator) end(_ context.Context, _ chan<- Item) { _ = "STUB: not implemented"; return }

func (op *takeOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// TakeLast emits only the last n items emitted by an Observable.
	// Cannot be run in parallel.
	return
}

func (o *ObservableImpl) TakeLast(nth uint, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type takeLast struct {
	n     int
	r     *ring.Ring
	count int
}

func (op *takeLast) next(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *takeLast) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *takeLast) end(ctx context.Context, dst chan<- Item) { _ = "STUB: not implemented"; return }

func (op *takeLast) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// TakeUntil returns an Observable that emits items emitted by the source Observable,
	// checks the specified predicate for each item, and then completes when the condition is satisfied.
	// Cannot be run in parallel.
	return
}

func (o *ObservableImpl) TakeUntil(apply Predicate, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type takeUntilOperator struct {
	apply Predicate
}

func (op *takeUntilOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *takeUntilOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *takeUntilOperator) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *takeUntilOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// TakeWhile returns an Observable that emits items emitted by the source ObservableSource so long as each
	// item satisfied a specified condition, and then completes as soon as this condition is not satisfied.
	// Cannot be run in parallel.
	return
}

func (o *ObservableImpl) TakeWhile(apply Predicate, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type takeWhileOperator struct {
	apply Predicate
}

func (op *takeWhileOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *takeWhileOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *takeWhileOperator) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *takeWhileOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// TimeInterval converts an Observable that emits items into one that emits indications of the amount of time elapsed between those emissions.
	return
}

func (o *ObservableImpl) TimeInterval(opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Timestamp attaches a timestamp to each item emitted by an Observable indicating when it was emitted.
func (o *ObservableImpl) Timestamp(opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type timestampOperator struct {
}

func (op *timestampOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *timestampOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *timestampOperator) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *timestampOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// ToMap convert the sequence of items emitted by an Observable
	// into a map keyed by a specified key function.
	// Cannot be run in parallel.
	return
}

func (o *ObservableImpl) ToMap(keySelector Func, opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type toMapOperator struct {
	keySelector Func
	m           map[interface{}]interface{}
}

func (op *toMapOperator) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *toMapOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *toMapOperator) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *toMapOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// ToMapWithValueSelector convert the sequence of items emitted by an Observable
	// into a map keyed by a specified key function and valued by another
	// value function.
	// Cannot be run in parallel.
	return
}

func (o *ObservableImpl) ToMapWithValueSelector(keySelector, valueSelector Func, opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

type toMapWithValueSelector struct {
	keySelector, valueSelector Func
	m                          map[interface{}]interface{}
}

func (op *toMapWithValueSelector) next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *toMapWithValueSelector) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *toMapWithValueSelector) end(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *toMapWithValueSelector) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// ToSlice collects all items from an Observable and emit them in a slice and an optional error.
	// Cannot be run in parallel.
	return
}

func (o *ObservableImpl) ToSlice(initialCapacity int, opts ...Option) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type toSliceOperator struct {
	s             []interface{}
	observableErr error
}

func (op *toSliceOperator) next(_ context.Context, item Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *toSliceOperator) err(_ context.Context, item Item, _ chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *toSliceOperator) end(_ context.Context, _ chan<- Item) { _ = "STUB: not implemented"; return }

func (op *toSliceOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// Unmarshal transforms the items emitted by an Observable by applying an unmarshalling to each item.
	return
}

func (o *ObservableImpl) Unmarshal(unmarshaller Unmarshaller, factory func() interface{}, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// WindowWithCount periodically subdivides items from an Observable into Observable windows of a given size and emit these windows
// rather than emitting the items one at a time.
func (o *ObservableImpl) WindowWithCount(count int, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type windowWithCountOperator struct {
	count          int
	iCount         int
	currentChannel chan Item
	option         Option
}

func (op *windowWithCountOperator) pre(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *windowWithCountOperator) post(ctx context.Context, dst chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *windowWithCountOperator) next(ctx context.Context, item Item, dst chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *windowWithCountOperator) err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func (op *windowWithCountOperator) end(_ context.Context, _ chan<- Item) {
	_ = "STUB: not implemented"
	return
}

func (op *windowWithCountOperator) gatherNext(_ context.Context, _ Item, _ chan<- Item, _ operatorOptions) {
	_ = "STUB: not implemented"

	// WindowWithTime periodically subdivides items from an Observable into Observables based on timed windows
	// and emit them rather than emitting the items one at a time.
	return
}

func (o *ObservableImpl) WindowWithTime(timespan Duration, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// WindowWithTimeOrCount periodically subdivides items from an Observable into Observables based on timed windows or a specific size
// and emit them rather than emitting the items one at a time.
func (o *ObservableImpl) WindowWithTimeOrCount(timespan Duration, count int, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// ZipFromIterable merges the emissions of an Iterable via a specified function
// and emit single items for each combination based on the results of this function.
func (o *ObservableImpl) ZipFromIterable(iterable Iterable, zipper Func2, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}
