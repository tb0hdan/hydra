package hydra

import "context"

type WorkerNode interface {
	Process(ctx context.Context, item interface{}) (interface{}, error)
	GetItem(ctx context.Context) (interface{}, error)
	SubmitResult(ctx context.Context, result interface{}) error
}
