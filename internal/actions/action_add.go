package actions

import (
	"fmt"
	"todo/internal/context"
	"todo/internal/model"
)

func init() {
	RegAction("add", func(ctx *context.Context) error {
		c, err := ctx.Count(ctx)
		if err != nil {
			return err
		}
		keys := ctx.Keys()
		if len(keys) != 1 {
			return fmt.Errorf("参数错误： %v", keys)
		}
		todo := model.Todo{Status: model.Undo, Index: c + 1, Content: keys[0]}
		ctx.WriteLine(fmt.Sprintf("1. %v\n", todo.Content))
		ctx.WriteLine(fmt.Sprintf("Item %v added", todo.Index))
		return ctx.Create(ctx, todo)
	})
}
