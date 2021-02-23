package actions

import (
	"fmt"
	"todo/internal/context"
	"todo/internal/dao"
	todo2 "todo/internal/model/todo"
)

func init() {
	RegAction("list", func(ctx *context.Context) error {
		var queries []todo2.QueryFunc
		if !ctx.Params.Bool("all") {
			queries = append(queries, todo2.QueryStatus(todo2.Undo))
		}
		todos, err := dao.Todo().List(ctx, queries...)
		if err != nil {
			return err
		}
		ctx.WriteToOutput("\n")
		for _, todo := range todos {
			var output string
			if todo.Status == todo2.Done {
				output = fmt.Sprintf("%v. [%v] %v", todo.Index, todo.Status, todo.Content)
			} else {
				output = fmt.Sprintf("%v. %v", todo.Index, todo.Content)
			}
			ctx.WriteToOutput(output)
		}
		ctx.WriteToOutput("\n")
		ctx.WriteToOutput(fmt.Sprintf("Total: %v items", len(todos)))
		return nil
	})
}
