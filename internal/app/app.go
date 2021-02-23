/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:17
 */

package app

import (
	context2 "context"
	"fmt"
	"io"
	"todo/internal/actions"
	"todo/internal/context"
)

type App struct {
	parse context.ParseParam
}

func NewApp(param context.ParseParam) *App {
	return &App{parse: param}
}

func (e App) Run(action string, input io.Reader, output io.Writer) {
	//	 初始化上下文
	p := context.Params{}
	e.parse(input, p)

	ctx := &context.Context{
		Context:  context2.Background(),
		Request:  context.Request{Params: p},
		Response: &context.Response{Output: output},
	}

	if err := actions.DoAction(action, ctx); err != nil {
		fmt.Println(err)
	}
}
