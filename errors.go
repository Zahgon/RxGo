package rxgo

// IllegalInputError is triggered when the observable receives an illegal input.
type IllegalInputError struct {
	error string
}

func (e IllegalInputError) Error() string { _ = "STUB: not implemented"; return "" }

// IndexOutOfBoundError is triggered when the observable cannot access to the specified index.
type IndexOutOfBoundError struct {
	error string
}

func (e IndexOutOfBoundError) Error() string { _ = "STUB: not implemented"; return "" }
