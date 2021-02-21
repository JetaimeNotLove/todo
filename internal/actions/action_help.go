package actions

import (
	"todo/internal/context"
)

func init() {
	RegAction("help", func(ctx *context.Context) error {
		return nil
	})
}
