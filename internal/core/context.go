/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午2:56
 */

package core

import (
	"io"
	"strconv"
)

type Context struct {
	Store TodoStore
	Req   Request
	Resp  *Response
}

type Param map[string]string

type ParseParam func(io.Reader, Param)

func (p Param) Add(key, val string) {
	p[key] = val
}

func (p Param) String(key string) string {
	return p[key]
}

func (p Param) Int(key string) int {
	val, ok := p[key]
	if !ok {
		return 0
	}
	res, _ := strconv.Atoi(val)
	return res
}

func (p Param) Bool(key string) bool {
	_, ok := p[key]
	return ok
}

type Request struct {
	Param
}

type Response struct {
	output io.Writer
}
