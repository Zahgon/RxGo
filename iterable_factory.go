package rxgo

type factoryIterable struct {
	factory func(opts ...Option) <-chan Item
}

func newFactoryIterable(factory func(opts ...Option) <-chan Item) Iterable {
	_ = "STUB: not implemented"
	return *new(Iterable)
}

func (i *factoryIterable) Observe(opts ...Option) <-chan Item {
	_ = "STUB: not implemented"
	return nil
}
