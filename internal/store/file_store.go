/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:09
 */

package store

import (
	"todo/internal/core"
)

type FileStore struct {
}

func (f FileStore) List(ctx core.Context, queries ...core.QueryFunc) (core.Todos, error) {
	panic("implement me")
}

func (f FileStore) Create(ctx core.Context, todo core.Todo) error {
	panic("implement me")
}

func (f FileStore) Update(ctx core.Context, index int, updates ...core.UpdateFunc) error {
	panic("implement me")
}
