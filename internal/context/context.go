/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午2:56
 */

package context

import (
	"io"
	"todo/internal/store"
)

type Context struct {
	Store store.TodoStore
	Req   Request
	Resp  *Response
}

type Response struct {
	Output io.Writer
}
