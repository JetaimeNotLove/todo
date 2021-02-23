package actions

import (
	"fmt"
	"todo/internal/context"
	"todo/internal/dao"
	"todo/internal/model"
)

func init() {
	RegAction("list", func(ctx *context.Context) error {
		var queries []dao.QueryFunc
		if !ctx.Params.Bool("all") {
			queries = append(queries, dao.QueryStatus(model.Undo))
		}
		todos, err := ctx.Todo.List(ctx, queries...)
		if err != nil {
			return err
		}
		ctx.WriteToOutput("\n")
		for _, todo := range todos {
			var output string
			if todo.Status == model.Done {
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
