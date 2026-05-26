// Package rxgo is the main RxGo package.
package rxgo

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
)

// Observable is the standard interface for Observables.
type Observable interface {
	Iterable
	All(predicate Predicate, opts ...Option) Single
	AverageFloat32(opts ...Option) Single
	AverageFloat64(opts ...Option) Single
	AverageInt(opts ...Option) Single
	AverageInt8(opts ...Option) Single
	AverageInt16(opts ...Option) Single
	AverageInt32(opts ...Option) Single
	AverageInt64(opts ...Option) Single
	BackOffRetry(backOffCfg backoff.BackOff, opts ...Option) Observable
	BufferWithCount(count int, opts ...Option) Observable
	BufferWithTime(timespan Duration, opts ...Option) Observable
	BufferWithTimeOrCount(timespan Duration, count int, opts ...Option) Observable
	Connect(ctx context.Context) (context.Context, Disposable)
	Contains(equal Predicate, opts ...Option) Single
	Count(opts ...Option) Single
	Debounce(timespan Duration, opts ...Option) Observable
	DefaultIfEmpty(defaultValue interface{}, opts ...Option) Observable
	Distinct(apply Func, opts ...Option) Observable
	DistinctUntilChanged(apply Func, opts ...Option) Observable
	DoOnCompleted(completedFunc CompletedFunc, opts ...Option) Disposed
	DoOnError(errFunc ErrFunc, opts ...Option) Disposed
	DoOnNext(nextFunc NextFunc, opts ...Option) Disposed
	ElementAt(index uint, opts ...Option) Single
	Error(opts ...Option) error
	Errors(opts ...Option) []error
	Filter(apply Predicate, opts ...Option) Observable
	Find(find Predicate, opts ...Option) OptionalSingle
	First(opts ...Option) OptionalSingle
	FirstOrDefault(defaultValue interface{}, opts ...Option) Single
	FlatMap(apply ItemToObservable, opts ...Option) Observable
	ForEach(nextFunc NextFunc, errFunc ErrFunc, completedFunc CompletedFunc, opts ...Option) Disposed
	GroupBy(length int, distribution func(Item) int, opts ...Option) Observable
	GroupByDynamic(distribution func(Item) string, opts ...Option) Observable
	IgnoreElements(opts ...Option) Observable
	Join(joiner Func2, right Observable, timeExtractor func(interface{}) time.Time, window Duration, opts ...Option) Observable
	Last(opts ...Option) OptionalSingle
	LastOrDefault(defaultValue interface{}, opts ...Option) Single
	Map(apply Func, opts ...Option) Observable
	Marshal(marshaller Marshaller, opts ...Option) Observable
	Max(comparator Comparator, opts ...Option) OptionalSingle
	Min(comparator Comparator, opts ...Option) OptionalSingle
	OnErrorResumeNext(resumeSequence ErrorToObservable, opts ...Option) Observable
	OnErrorReturn(resumeFunc ErrorFunc, opts ...Option) Observable
	OnErrorReturnItem(resume interface{}, opts ...Option) Observable
	Reduce(apply Func2, opts ...Option) OptionalSingle
	Repeat(count int64, frequency Duration, opts ...Option) Observable
	Retry(count int, shouldRetry func(error) bool, opts ...Option) Observable
	Run(opts ...Option) Disposed
	Sample(iterable Iterable, opts ...Option) Observable
	Scan(apply Func2, opts ...Option) Observable
	SequenceEqual(iterable Iterable, opts ...Option) Single
	Send(output chan<- Item, opts ...Option)
	Serialize(from int, identifier func(interface{}) int, opts ...Option) Observable
	Skip(nth uint, opts ...Option) Observable
	SkipLast(nth uint, opts ...Option) Observable
	SkipWhile(apply Predicate, opts ...Option) Observable
	StartWith(iterable Iterable, opts ...Option) Observable
	SumFloat32(opts ...Option) OptionalSingle
	SumFloat64(opts ...Option) OptionalSingle
	SumInt64(opts ...Option) OptionalSingle
	Take(nth uint, opts ...Option) Observable
	TakeLast(nth uint, opts ...Option) Observable
	TakeUntil(apply Predicate, opts ...Option) Observable
	TakeWhile(apply Predicate, opts ...Option) Observable
	TimeInterval(opts ...Option) Observable
	Timestamp(opts ...Option) Observable
	ToMap(keySelector Func, opts ...Option) Single
	ToMapWithValueSelector(keySelector, valueSelector Func, opts ...Option) Single
	ToSlice(initialCapacity int, opts ...Option) ([]interface{}, error)
	Unmarshal(unmarshaller Unmarshaller, factory func() interface{}, opts ...Option) Observable
	WindowWithCount(count int, opts ...Option) Observable
	WindowWithTime(timespan Duration, opts ...Option) Observable
	WindowWithTimeOrCount(timespan Duration, count int, opts ...Option) Observable
	ZipFromIterable(iterable Iterable, zipper Func2, opts ...Option) Observable
}

// ObservableImpl implements Observable.
type ObservableImpl struct {
	parent   context.Context
	iterable Iterable
}

func defaultErrorFuncOperator(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions) {
	_ = "STUB: not implemented"
	return
}

func customObservableOperator(parent context.Context, f func(ctx context.Context, next chan Item, option Option, opts ...Option), opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

type operator interface {
	next(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions)
	err(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions)
	end(ctx context.Context, dst chan<- Item)
	gatherNext(ctx context.Context, item Item, dst chan<- Item, operatorOptions operatorOptions)
}

func observable(parent context.Context, iterable Iterable, operatorFactory func() operator, forceSeq, bypassGather bool, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

func single(parent context.Context, iterable Iterable, operatorFactory func() operator, forceSeq, bypassGather bool, opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

func optionalSingle(parent context.Context, iterable Iterable, operatorFactory func() operator, forceSeq, bypassGather bool, opts ...Option) OptionalSingle {
	_ = "STUB: not implemented"
	return *new(OptionalSingle)
}

func runSequential(ctx context.Context, next chan Item, iterable Iterable, operatorFactory func() operator, option Option, opts ...Option) {
	_ = "STUB: not implemented"
	return
}

func runParallel(ctx context.Context, next chan Item, observe <-chan Item, operatorFactory func() operator, bypassGather bool, option Option, opts ...Option) {
	_ = "STUB: not implemented"
	return
}

// Gather

// Scatter

func runFirstItem(ctx context.Context, f func(interface{}) int, notif chan Item, observe <-chan Item, next chan Item, operatorFactory func() operator, option Option, opts ...Option) {
	_ = "STUB: not implemented"
	return
}

func (o *ObservableImpl) serialize(parent context.Context, fromCh chan Item, identifier func(interface{}) int, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}
