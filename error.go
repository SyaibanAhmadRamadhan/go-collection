package collection

import (
	"errors"
	"fmt"
	"runtime"
)

type StackErr struct {
	stackErr error

	// opt using http status code
	httpCode int
}

func (s *StackErr) Error() string {
	return s.stackErr.Error()
}

func (s *StackErr) Unwrap() error {
	currentErr := s.stackErr
	for {
		err := errors.Unwrap(currentErr)
		if err == nil {
			return currentErr
		}

		currentErr = err
	}
}

func (s *StackErr) GetHttpCode() int {
	return s.httpCode
}

type Option func(s *StackErr)

func SetHttpCode(code int) Option {
	return func(s *StackErr) {
		s.httpCode = code
	}
}

func Err(err error, opts ...Option) *StackErr {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	s := &StackErr{
		stackErr: fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err),
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
