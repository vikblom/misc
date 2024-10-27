package wrap

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type Error struct {
	inner error
	pc    uintptr
}

func Errorf(format string, args ...any) error {
	s := fmt.Errorf(format, args...)

	var pc [1]uintptr
	n := runtime.Callers(2, pc[:])
	if n != 1 {
		panic("no caller")
	}
	return Error{
		inner: s,
		pc:    pc[0],
	}
}

func (e Error) Error() string {
	frames := runtime.CallersFrames([]uintptr{e.pc})
	f, _ := frames.Next()
	return fmt.Sprintf("(%s:%d) ", filepath.Base(f.File), f.Line) + e.inner.Error()
}

func (e Error) Unwrap() error {
	return e.inner
}
