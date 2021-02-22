package actions

import (
	"fmt"
	"todo/internal/context"
)

func init() {
	RegAction("list", func(ctx *context.Context) error {
		todos, err := ctx.Todo.List(ctx)
		if err != nil {
			return err
		}
		ctx.WriteToOutput("\n")
		for _, todo := range todos {
			ctx.WriteToOutput(fmt.Sprintf("%v. %v", todo.Index, todo.Content))
		}
		ctx.WriteToOutput("\n")
		ctx.WriteToOutput(fmt.Sprintf("Total: %v items", len(todos)))
		return nil
	})
}
