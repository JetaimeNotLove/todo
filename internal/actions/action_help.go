package actions

import (
	"todo/internal/context"
)

func init() {
	RegAction("help", func(ctx *context.Context) error {
		ctx.WriteToOutput("list")
		ctx.WriteToOutput("add")
		ctx.WriteToOutput("done")
		return nil
	})
}
