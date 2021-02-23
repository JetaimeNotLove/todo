package actions

import (
	"fmt"
	"strconv"
	"todo/internal/context"
	"todo/internal/dao"
	"todo/internal/model/todo"
)

func init() {
	RegAction("done", func(ctx *context.Context) error {
		keys := ctx.Params.Keys()
		if len(keys) != 1 {
			return fmt.Errorf("参数错误： %v", keys)
		}
		index, err := strconv.Atoi(keys[0])
		if err != nil {
			return err
		}
		ctx.WriteToOutput(fmt.Sprintf("Item %v done.", index))
		return dao.Todo().Update(ctx, index, todo.UpdateStatus(todo.Done))
	})
}
