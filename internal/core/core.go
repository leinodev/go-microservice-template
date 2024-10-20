package core

import "context"

type Core[In, Out any] interface {
	Execute(ctx context.Context, input In) (Out, error)
}
