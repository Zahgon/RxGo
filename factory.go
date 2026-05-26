package rxgo

// Amb takes several Observables, emit all of the items from only the first of these Observables
// to emit an item or notification.
func Amb(observables []Observable, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// CombineLatest combines the latest item emitted by each Observable via a specified function
// and emit items based on the results of this function.
func CombineLatest(f FuncN, observables []Observable, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Concat emits the emissions from two or more Observables without interleaving them.
func Concat(observables []Observable, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Create creates an Observable from scratch by calling observer methods programmatically.
func Create(f []Producer, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Defer does not create the Observable until the observer subscribes,
// and creates a fresh Observable for each observer.
func Defer(f []Producer, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Empty creates an Observable with no item and terminate immediately.
func Empty() Observable { _ = "STUB: not implemented"; return *new(Observable) }

// FromChannel creates a cold observable from a channel.
func FromChannel(next <-chan Item, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// FromEventSource creates a hot observable from a channel.
func FromEventSource(next <-chan Item, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Interval creates an Observable emitting incremental integers infinitely between
// each given time interval.
func Interval(interval Duration, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Just creates an Observable with the provided items.
func Just(items ...interface{}) func(opts ...Option) Observable {
	_ = "STUB: not implemented"
	return nil
}

// JustItem creates a single from one item.
func JustItem(item interface{}, opts ...Option) Single {
	_ = "STUB: not implemented"
	return *new(Single)
}

// Merge combines multiple Observables into one by merging their emissions
func Merge(observables []Observable, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Never creates an Observable that emits no items and does not terminate.
func Never() Observable { _ = "STUB: not implemented"; return *new(Observable) }

// Range creates an Observable that emits count sequential integers beginning
// at start.
func Range(start, count int, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Start creates an Observable from one or more directive-like Supplier
// and emits the result of each operation asynchronously on a new Observable.
func Start(fs []Supplier, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}

// Thrown creates an Observable that emits no items and terminates with an error.
func Thrown(err error) Observable { _ = "STUB: not implemented"; return *new(Observable) }

// Timer returns an Observable that completes after a specified delay.
func Timer(d Duration, opts ...Option) Observable {
	_ = "STUB: not implemented"
	return *new(Observable)
}
