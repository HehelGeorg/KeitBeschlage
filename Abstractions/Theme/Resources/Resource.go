package Resources

import "context"

type IResource interface {
	ApplyResource(context context.Context) error
}
