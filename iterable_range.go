package rxgo

type rangeIterable struct {
	start, count int
	opts         []Option
}

func newRangeIterable(start, count int, opts ...Option) Iterable {
	_ = "STUB: not implemented"
	return *new(Iterable)
}

func (i *rangeIterable) Observe(opts ...Option) <-chan Item { _ = "STUB: not implemented"; return nil }
