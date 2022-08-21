package xio

import (
	"errors"
	"testing"

	"go.uber.org/multierr"
)

type mockCloser struct {
	err error
}

func (m mockCloser) Close() error {
	return m.err
}

func TestMultiCloser(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		got := MultiCloser(
			mockCloser{
				err: nil,
			},
			mockCloser{
				err: nil,
			},
		).Close()
		if got != nil {
			t.Errorf("unexpected err: %s", got)
			return
		}
	})
	t.Run("single err", func(t *testing.T) {
		want := errors.New("failed")
		got := MultiCloser(
			mockCloser{
				err: nil,
			},
			mockCloser{
				err: want,
			},
		).Close()
		if !errors.Is(got, want) {
			t.Errorf("mismatch err: got: '%s', want: '%s'", got, want)
			return
		}
	})
	t.Run("multi err", func(t *testing.T) {
		failed1Err := errors.New("failed1")
		failed2Err := errors.New("failed2")
		want := multierr.Combine(failed1Err, failed2Err)
		got := MultiCloser(
			mockCloser{
				err: failed1Err,
			},
			mockCloser{
				err: failed2Err,
			},
		).Close()
		if got.Error() != want.Error() {
			t.Errorf("mismatch err: got: '%s', want: '%s'", got, want)
			return
		}
	})

}
