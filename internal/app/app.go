/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:17
 */

package app

import (
	context2 "context"
	"io"
	"todo/internal/actions"
	"todo/internal/context"
	store2 "todo/internal/store"
)

type App struct {
	parse context.ParseParam
	store store2.TodoStore
}

func NewEngine(param context.ParseParam, store store2.TodoStore) *App {
	return &App{parse: param, store: store}
}

func (e App) Run(action string, param io.Reader, writer io.Writer) {
	//	 初始化上下文
	p := context.Params{}
	e.parse(param, p)

	ctx := &context.Context{
		context2.Background(),
		e.store,
		context.Request{Params: p},
		&context.Response{Output: writer},
	}

	if err := actions.DoAction(action, ctx); err != nil {
		panic(err)
	}
}
