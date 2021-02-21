/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午2:56
 */

package context

import (
	"context"
	"io"
	"todo/internal/store"
)

type Context struct {
	context.Context
	store.TodoStore
	Request
	*Response
}

func (c *Context) WriteLine(s string) {
	_, _ = c.Output.Write([]byte(s))
	_, _ = c.Output.Write([]byte{'\n'})
}

type Response struct {
	Output io.Writer
}
