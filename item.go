package rxgo

import (
	"context"
	"time"
)

type (
	// Item is a wrapper having either a value or an error.
	Item struct {
		V interface{}
		E error
	}

	// TimestampItem attach a timestamp to an item.
	TimestampItem struct {
		Timestamp time.Time
		V         interface{}
	}

	// CloseChannelStrategy indicates a strategy on whether to close a channel.
	CloseChannelStrategy uint32
)

const (
	// LeaveChannelOpen indicates to leave the channel open after completion.
	LeaveChannelOpen CloseChannelStrategy = iota
	// CloseChannel indicates to close the channel open after completion.
	CloseChannel
)

// Of creates an item from a value.
func Of(i interface{}) Item {
	_ = "STUB: not implemented"

	// Error creates an item from an error.
	return *new(Item)
}

func Error(err error) Item {
	_ = "STUB: not implemented"
	return *

	// SendItems is an utility function that send a list of interface{} and indicate a strategy on whether to close
	// the channel once the function completes.
	new(Item)
}

func SendItems(ctx context.Context, ch chan<- Item, strategy CloseChannelStrategy, items ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func send(ctx context.Context, ch chan<- Item, items ...interface{}) {
	_ = "STUB: not implemented"
	return
}

// Error checks if an item is an error.
func (i Item) Error() bool {
	_ = "STUB: not implemented"

	// SendBlocking sends an item and blocks until it is sent.
	return false
}

func (i Item) SendBlocking(ch chan<- Item) {
	_ = "STUB: not implemented"

	// SendContext sends an item and blocks until it is sent or a context canceled.
	// It returns a boolean to indicate whether the item was sent.
	return
}

func (i Item) SendContext(ctx context.Context, ch chan<- Item) bool {
	_ = "STUB: not implemented"
	return false

	// Context's done channel has the highest priority
}

// SendNonBlocking sends an item without blocking.
// It returns a boolean to indicate whether the item was sent.
func (i Item) SendNonBlocking(ch chan<- Item) bool { _ = "STUB: not implemented"; return false }
