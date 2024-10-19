package collection

import "context"

type CloseFn func() (err error)
type CloseFnCtx func(ctx context.Context) (err error)
