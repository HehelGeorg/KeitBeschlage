package UnixManager

import (
	"context"
	"sync"
)

type UnixManager interface {
	Run(ctx context.Context, wg *sync.WaitGroup) error
}
