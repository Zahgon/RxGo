package rxgo

type justIterable struct {
	items []interface{}
	opts  []Option
}

func newJustIterable(items ...interface{}) func(opts ...Option) Iterable {
	_ = "STUB: not implemented"
	return nil
}

func (i *justIterable) Observe(opts ...Option) <-chan Item { _ = "STUB: not implemented"; return nil }
