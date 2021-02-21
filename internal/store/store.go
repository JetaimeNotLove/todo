/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:00
 */

package store

import (
	"context"
	"todo/internal/model"
)

type QueryParam struct {
	Status model.Status
}

type QueryFunc func(param *QueryParam)

type UpdateParam struct {
	M map[string]interface{}
}

type UpdateFunc func(param *UpdateParam)

type TodoStore interface {
	List(ctx context.Context, queries ...QueryFunc) (model.Todos, error)
	Create(ctx context.Context, todo model.Todo) error
	Update(ctx context.Context, index int, updates ...UpdateFunc) error
	Count(ctx context.Context) (int, error)
}
