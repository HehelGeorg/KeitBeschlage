package GrpcServer

import "context"

type GrpcServer interface {
	Up(ctx context.Context) error
}
