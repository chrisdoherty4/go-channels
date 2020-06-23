type Context interface {
    // Deadline returns the time when work done on behalf of this context
    // should be canceled. Deadline returns ok==false when no deadline is
    // set. Successive calls to Deadline return the same results.
    Deadline() (deadline time.Time, ok bool)

    // Done returns a channel that's closed when work done on behalf of this
    // context should be canceled. Done may return nil if this context can
    // never be canceled. Successive calls to Done return the same value.
    // The close of the Done channel may happen asynchronously,
    // after the cancel function returns.
    Done() <-chan struct{}

    // If Done is not yet closed, Err returns nil.
    // If Done is closed, Err returns a non-nil error explaining why:
    // Canceled if the context was canceled
    // or DeadlineExceeded if the context's deadline passed.
    // After Err returns a non-nil error, successive calls to Err return the same error.
    Err() error

    // Value returns the value associated with this context for key, or nil
    // if no value is associated with key. Successive calls to Value with
    // the same key returns the same result.
    Value(key interface{}) interface{}
}