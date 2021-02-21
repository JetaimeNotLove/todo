/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:17
 */

package core

import "io"

type Engine struct {
	parse ParseParam
	store TodoStore
}

func NewEngine(param ParseParam, store TodoStore) *Engine {
	return &Engine{parse: param, store: store}
}

func (e Engine) Run(action string, param io.Reader, writer io.Writer) {
	//	 初始化上下文
	p := Param{}
	e.parse(param, p)

	ctx := &Context{
		Store: e.store,
		Req:   Request{p},
		Resp:  &Response{writer},
	}

	DoAction(action, ctx)
}
