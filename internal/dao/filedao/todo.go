/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:09
 */

package filedao

import (
	"context"
	"errors"
	"reflect"
	"todo/internal/model/todo"
)

type TodoDao struct {
}

func (fs *TodoDao) List(ctx context.Context, queries ...todo.QueryFunc) (todo.Todos, error) {
	todos := todo.Todos{}
	if err := read(todolist, &todos); err != nil {
		return nil, err
	}
	filters := []todo.FilterFunc{}
	qp := new(todo.QueryParam)
	for _, query := range queries {
		query(qp)
	}
	if qp.Status != 0 {
		filters = append(filters, func(todo todo.Todo) bool {
			return todo.Status == qp.Status
		})
	}
	return todos.Filter(todo.MergeFilter(filters...)), nil
}

func (fs *TodoDao) Create(ctx context.Context, todo todo.Todo) error {
	todos, err := fs.List(ctx)
	if err != nil {
		return err
	}
	todos.Add(todo)
	return write(todolist, &todos)

}
func (fs *TodoDao) Update(ctx context.Context, index int, updates ...todo.UpdateFunc) error {
	todos, err := fs.List(ctx)
	if err != nil {
		return err
	}
	todoTiem := todos.Get(index)
	if todoTiem == nil {
		return errors.New("不存在todo")
	}
	up := todo.UpdateParam{M: map[string]interface{}{}}
	for _, update := range updates {
		update(&up)
	}
	v := reflect.ValueOf(todoTiem).Elem()
	for key, val := range up.M {
		f := v.FieldByName(key)
		if !f.IsZero() {
			f.Set(reflect.ValueOf(val))
		}
	}
	return write(todolist, &todos)
}
func (fs *TodoDao) Count(ctx context.Context) (int, error) {
	todos, err := fs.List(ctx)
	if err != nil {
		return 0, err
	}
	return todos.Size(), nil
}
