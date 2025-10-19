package SocketLIstener

import (
	"context"
)

type SocketListener interface {
	Listen(ctx context.Context) (<-chan int, error)
}
