/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:00
 */

package dao

import (
	"context"
	"todo/internal/model/todo"
)

var todoDao TodoDao

func SetTodo(dao TodoDao) {
	todoDao = dao
}
func Todo() TodoDao {
	return todoDao
}

type TodoDao interface {
	List(ctx context.Context, queries ...todo.QueryFunc) (todo.Todos, error)
	Create(ctx context.Context, todo todo.Todo) error
	Update(ctx context.Context, index int, updates ...todo.UpdateFunc) error
	Count(ctx context.Context) (int, error)
}
