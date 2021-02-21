/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:09
 */

package store

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"todo/internal/model"
	"todo/pkg/file"
)

const (
	dir      = "./.todolist"
	todolist = "todolist.json"
)

type FileStore struct {
}

func (fs *FileStore) Init() {
	if !file.Exists(dir) {
		if err := os.Mkdir(dir, os.ModePerm|os.ModeDir); err != nil {
			panic(err)
		}
	}
	if !file.Exists(filepath.Join(dir, todolist)) {
		f, err := os.Create(filepath.Join(dir, todolist))
		if err != nil {
			panic(f)
		}
		f.Close()
	}
}

func (fs *FileStore) List(ctx context.Context, queries ...QueryFunc) (model.Todos, error) {
	fileName := filepath.Join(dir, todolist)
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	todos := model.Todos{}
	if len(data) == 0 {
		return todos, nil
	}
	if err := json.Unmarshal(data, &todos); err != nil {
		return nil, err
	}
	filters := []model.FilterFunc{}
	qp := new(QueryParam)
	for _, query := range queries {
		query(qp)
	}
	if qp.Status != 0 {
		filters = append(filters, func(todo model.Todo) bool {
			return todo.Status == qp.Status
		})
	}
	return todos.Filter(model.MergeFilter(filters...)), nil
}

func (fs *FileStore) Create(ctx context.Context, todo model.Todo) error {
	todos, err := fs.List(ctx)
	if err != nil {
		return err
	}
	todos.Add(todo)
	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(dir, todolist), data, os.ModePerm)

}
func (fs *FileStore) Update(ctx context.Context, index int, updates ...UpdateFunc) error {
	panic("implement me")
}
func (fs *FileStore) Count(ctx context.Context) (int, error) {
	todos, err := fs.List(ctx)
	if err != nil {
		return 0, err
	}
	return todos.Size(), nil
}
