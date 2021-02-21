package actions

import (
	"fmt"
	"todo/internal/context"
)

func init() {
	RegAction("list", func(ctx *context.Context) error {
		todos, err := ctx.List(ctx)
		if err != nil {
			return err
		}
		ctx.WriteLine("\n")
		for _, todo := range todos {
			ctx.WriteLine(fmt.Sprintf("%v. %v", todo.Index, todo.Content))
		}
		ctx.WriteLine("\n")
		ctx.WriteLine(fmt.Sprintf("Total: %v items", len(todos)))
		return nil
	})
}
