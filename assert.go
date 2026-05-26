package rxgo

import (
	"context"
	"testing"
)

// AssertPredicate is a custom predicate based on the items.
type AssertPredicate func(items []interface{}) error

// RxAssert lists the Observable assertions.
type RxAssert interface {
	apply(*rxAssert)
	itemsToBeChecked() (bool, []interface{})
	itemsNoOrderedToBeChecked() (bool, []interface{})
	noItemsToBeChecked() bool
	someItemsToBeChecked() bool
	raisedErrorToBeChecked() (bool, error)
	raisedErrorsToBeChecked() (bool, []error)
	raisedAnErrorToBeChecked() (bool, error)
	notRaisedErrorToBeChecked() bool
	itemToBeChecked() (bool, interface{})
	noItemToBeChecked() (bool, interface{})
	customPredicatesToBeChecked() (bool, []AssertPredicate)
}

type rxAssert struct {
	f                       func(*rxAssert)
	checkHasItems           bool
	checkHasNoItems         bool
	checkHasSomeItems       bool
	items                   []interface{}
	checkHasItemsNoOrder    bool
	itemsNoOrder            []interface{}
	checkHasRaisedError     bool
	err                     error
	checkHasRaisedErrors    bool
	errs                    []error
	checkHasRaisedAnError   bool
	checkHasNotRaisedError  bool
	checkHasItem            bool
	item                    interface{}
	checkHasNoItem          bool
	checkHasCustomPredicate bool
	customPredicates        []AssertPredicate
}

func (ass *rxAssert) apply(do *rxAssert) { _ = "STUB: not implemented"; return }

func (ass *rxAssert) itemsToBeChecked() (bool, []interface{}) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ass *rxAssert) itemsNoOrderedToBeChecked() (bool, []interface{}) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ass *rxAssert) noItemsToBeChecked() bool { _ = "STUB: not implemented"; return false }

func (ass *rxAssert) someItemsToBeChecked() bool { _ = "STUB: not implemented"; return false }

func (ass *rxAssert) raisedErrorToBeChecked() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ass *rxAssert) raisedErrorsToBeChecked() (bool, []error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ass *rxAssert) raisedAnErrorToBeChecked() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ass *rxAssert) notRaisedErrorToBeChecked() bool { _ = "STUB: not implemented"; return false }

func (ass *rxAssert) itemToBeChecked() (bool, interface{}) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ass *rxAssert) noItemToBeChecked() (bool, interface{}) {
	_ = "STUB: not implemented"
	return false, nil
}

func (ass *rxAssert) customPredicatesToBeChecked() (bool, []AssertPredicate) {
	_ = "STUB: not implemented"
	return false, nil
}

func newAssertion(f func(*rxAssert)) *rxAssert { _ = "STUB: not implemented"; return nil }

// HasItems checks that the observable produces the corresponding items.
func HasItems(items ...interface{}) RxAssert { _ = "STUB: not implemented"; return *new(RxAssert) }

// HasItem checks if a single or optional single has a specific item.
func HasItem(i interface{}) RxAssert { _ = "STUB: not implemented"; return *new(RxAssert) }

// HasItemsNoOrder checks that an observable produces the corresponding items regardless of the order.
func HasItemsNoOrder(items ...interface{}) RxAssert {
	_ = "STUB: not implemented"
	return *new(RxAssert)
}

// IsNotEmpty checks that the observable produces some items.
func IsNotEmpty() RxAssert { _ = "STUB: not implemented"; return *new(RxAssert) }

// IsEmpty checks that the observable has not produce any item.
func IsEmpty() RxAssert { _ = "STUB: not implemented"; return *new(RxAssert) }

// HasError checks that the observable has produce a specific error.
func HasError(err error) RxAssert { _ = "STUB: not implemented"; return *new(RxAssert) }

// HasAnError checks that the observable has produce an error.
func HasAnError() RxAssert { _ = "STUB: not implemented"; return *new(RxAssert) }

// HasErrors checks that the observable has produce a set of errors.
func HasErrors(errs ...error) RxAssert { _ = "STUB: not implemented"; return *new(RxAssert) }

// HasNoError checks that the observable has not raised any error.
func HasNoError() RxAssert { _ = "STUB: not implemented"; return *new(RxAssert) }

// CustomPredicate checks a custom predicate.
func CustomPredicate(predicate AssertPredicate) RxAssert {
	_ = "STUB: not implemented"
	return *new(RxAssert)
}

func parseAssertions(assertions ...RxAssert) RxAssert {
	_ = "STUB: not implemented"
	return *new(RxAssert)
}

// Assert asserts the result of an iterable against a list of assertions.
func Assert(ctx context.Context, t *testing.T, iterable Iterable, assertions ...RxAssert) {
	_ = "STUB: not implemented"
	return
}
