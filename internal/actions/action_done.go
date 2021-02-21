package actions

import (
	"todo/internal/context"
)

func init() {
	RegAction("done", func(ctx *context.Context) error {
		return nil
	})
}
