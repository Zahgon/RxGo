package rxgo

type sliceIterable struct {
	items []Item
	opts  []Option
}

func newSliceIterable(items []Item, opts ...Option) Iterable {
	_ = "STUB: not implemented"
	return *new(Iterable)
}

func (i *sliceIterable) Observe(opts ...Option) <-chan Item { _ = "STUB: not implemented"; return nil }
