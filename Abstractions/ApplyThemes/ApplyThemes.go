package ApplyThemes

import (
	"context"
)

type ApplyThemes interface {
	Run(ctx context.Context) error
}
