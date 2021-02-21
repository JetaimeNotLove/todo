/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:09
 */

package store

import (
	"todo/internal/context"
	"todo/internal/model"
)

type FileStore struct {
}

func (f FileStore) Init() {

}

func (f FileStore) List(ctx context.Context, queries ...QueryFunc) (model.Todos, error) {
	qp := new(QueryParam)
	for _, query := range queries {
		query(qp)
	}
	if qp.Status != 0 {

	}
	return nil, nil
}

func (f FileStore) Create(ctx context.Context, todo model.Todo) error {
	panic("implement me")
}

func (f FileStore) Update(ctx context.Context, index int, updates ...UpdateFunc) error {
	panic("implement me")
}
