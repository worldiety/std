package std

import (
	"context"
)

var background = &Context{Value: context.Background()}

// Background returns a non-cancelable background context
func Background() *Context {
	return background
}

// WithCancel creates a sub context of parent which can be cancelled
func WithCancel(parent *Context) *Context {
	ctx, cancelFunc := context.WithCancel(parent.Value)
	return &Context{ctx, cancelFunc}
}

// WithValue creates a sub context of parent with the key and value pair. If the parent is cancelable, the child
// is not, so you have to wrap it again using WithCancel, otherwise if inheriting the cancellation from the parent
// the siblings would be also cancelled unexpectedly.
func WithValue(parent *Context, key *Box, value *Box) *Context {
	return &Context{context.WithValue(parent.Value, key.Unbox(), value.Unbox()), nil}
}

// WithDeadline returns a copy of the parent context with the deadline adjusted to be no later than d.
// If the parent's deadline is already earlier than d, WithDeadline(parent, d) is semantically equivalent
// to parent. The returned context's Done channel is closed when the deadline expires, when the returned
// cancel function is called, or when the parent context's Done channel is closed, whichever happens first.
// Canceling this context releases resources associated with it, so code should call cancel as soon as
// the operations running in this Context complete.
func WithDeadline(parent *Context, d *Time) *Context {
	ctx, f := context.WithDeadline(parent.Value, d.Value)
	return &Context{ctx, f}
}

// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
func WithTimeout(parent *Context, d *Duration) *Context {
	ctx, f := context.WithTimeout(parent.Value, d.Value)
	return &Context{ctx, f}
}

// A Context is a wrapper for the go standard library context
type Context struct {
	Value      context.Context // Value is never nil, discarded by gomobile
	CancelFunc func()          // CancelFunc can be nil. Use #Cancel(), discarded by gomobile
}

// Cancel tries to cancel the current context. Has no effect when called on contexts which cannot be cancelled.
// See also #WithDeadline(), #WithTimeout() and #WithCancel()
func (c *Context) Cancel() {
	f := c.CancelFunc
	if f != nil {
		f()
	}
}

// Unbox either returns #context.Background() or the boxed Context. Discarded by gomobile.
func (c *Context) Unbox() context.Context {
	if c == nil || c.Value == nil {
		return context.Background()
	}
	return c.Value
}

// Box returns this Context as a box. The cancel func will be lost.
func (c *Context) Box() *Box {
	return &Box{c.Value}
}
