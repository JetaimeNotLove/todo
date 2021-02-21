/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:17
 */

package app

import (
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
	p := context.Param{}
	e.parse(param, p)

	ctx := &context.Context{
		Store: e.store,
		Req:   context.Request{p},
		Resp:  &context.Response{Output: writer},
	}

	actions.DoAction(action, ctx)
}
