package rxgo

type deferIterable struct {
	fs   []Producer
	opts []Option
}

func newDeferIterable(f []Producer, opts ...Option) Iterable {
	_ = "STUB: not implemented"
	return *new(Iterable)
}

func (i *deferIterable) Observe(opts ...Option) <-chan Item { _ = "STUB: not implemented"; return nil }
