package Theme

import "context"

type ITheme interface {
	ApplyTheme(context context.Context) error
}
