package actions

import (
	"fmt"
	"todo/internal/context"
	"todo/internal/model"
)

func init() {
	RegAction("add", func(ctx *context.Context) error {
		c, err := ctx.Todo.Count(ctx)
		if err != nil {
			return err
		}
		keys := ctx.Params.Keys()
		if len(keys) != 1 {
			return fmt.Errorf("参数错误： %v", keys)
		}
		todo := model.Todo{Status: model.Undo, Index: c + 1, Content: keys[0]}
		ctx.WriteToOutput(fmt.Sprintf("1. %v\n", todo.Content))
		ctx.WriteToOutput(fmt.Sprintf("Item %v added", todo.Index))
		return ctx.Todo.Create(ctx, todo)
	})
}
