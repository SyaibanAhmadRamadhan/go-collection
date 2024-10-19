package collection

import (
	"fmt"
	"runtime"
)

type StackErr struct {
	stackErr error
	err      error

	// opt using http status code
	code int64
}

func (s *StackErr) Error() string {
	return s.stackErr.Error()
}

func (s *StackErr) Unwrap() error {
	return s.err
}

type Option func(s *StackErr)

func SetHttpCode(code int64) Option {
	return func(s *StackErr) {
		s.code = code
	}
}

func Err(err error, opts ...Option) *StackErr {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	s := &StackErr{
		stackErr: fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
		err:      err,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
