package actions

import (
	"fmt"
	"todo/internal/context"
	"todo/internal/dao"
	todo2 "todo/internal/model/todo"
)

func init() {
	RegAction("add", func(ctx *context.Context) error {
		c, err := dao.Todo().Count(ctx)
		if err != nil {
			return err
		}
		keys := ctx.Params.Keys()
		if len(keys) != 1 {
			return fmt.Errorf("参数错误： %v", keys)
		}
		todo := todo2.Todo{Status: todo2.Undo, Index: c + 1, Content: keys[0]}
		ctx.WriteToOutput(fmt.Sprintf("1. %v\n", todo.Content))
		ctx.WriteToOutput(fmt.Sprintf("Item %v added", todo.Index))
		return dao.Todo().Create(ctx, todo)
	})
}
