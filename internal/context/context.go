/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午2:56
 */

package context

import (
	"context"
	"io"
)

type Context struct {
	context.Context
	Request
	*Response

	value map[string]interface{}
}

func (c *Context) WriteToOutput(s string) {
	_, _ = c.Output.Write([]byte(s))
	_, _ = c.Output.Write([]byte{'\n'})
}

type Response struct {
	Output io.Writer
}
