package wrap_test

import (
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/vikblom/misc/wrap"
)

// Override the reported location to get stable test input.
//
//line foo.go:123
var err = wrap.Errorf("eof: %w", io.EOF)

func TestWrap(t *testing.T) {
	got := err

	want := "(foo.go:123) eof: EOF"
	if want != got.Error() {
		t.Errorf("\nwant: %q\ngot:  %q", want, got)
	}

	if !errors.Is(got, io.EOF) {
		t.Errorf("expected err to wrap io.EOF")
	}
}

var sink error

// go test -bench=. -count=10
func BenchmarkWrapf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := wrap.Errorf("foo")
		sink = err
	}
}

func BenchmarkFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := fmt.Errorf("foo")
		sink = err
	}
}
