package xio

import (
	"io"

	"go.uber.org/multierr"
)

type multiCloser struct {
	closers []io.Closer
}

func (mc multiCloser) Close() error {
	errs := make([]error, len(mc.closers))
	for i, closer := range mc.closers {
		errs[i] = closer.Close()
	}
	return multierr.Combine(errs...)
}

// MultiCloser creates a closer that iterates through all closers, returning all non-nil errors. 
func MultiCloser(closers ...io.Closer) io.Closer {
	return multiCloser{closers: closers}
}
