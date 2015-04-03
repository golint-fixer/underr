package underr_test

import (
	"errors"
	"testing"

	"github.com/daaku/underr"
	"github.com/facebookgo/ensure"
)

type wrapped struct {
	err error
}

func (w *wrapped) Error() string {
	return "wrapped:" + w.err.Error()
}

func (w *wrapped) Underlying() error {
	return w.err
}

func TestAllNil(t *testing.T) {
	ensure.True(t, underr.All(nil) == nil)
}

func TestAllDirect(t *testing.T) {
	e := errors.New("")
	ensure.DeepEqual(t, underr.All(e), []error{e})
}

func TestAllWrapped(t *testing.T) {
	u := errors.New("")
	w := &wrapped{err: u}
	ensure.DeepEqual(t, underr.All(w), []error{w, u})
}

func TestEqualsMatcherTrue(t *testing.T) {
	e := errors.New("")
	ensure.True(t, underr.Has(e, underr.Equal(e)))
}

func TestEqualsMatcherFalse(t *testing.T) {
	e := errors.New("")
	ensure.False(t, underr.Has(e, underr.Equal(nil)))
}
