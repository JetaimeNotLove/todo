/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午5:11
 */

package context

import (
	"io"
	"strconv"
)

type Params map[string]string

type ParseParam func(io.Reader, Params)

func (p Params) Add(key, val string) {
	p[key] = val
}

func (p Params) String(key string) string {
	return p[key]
}

func (p Params) Int(key string) int {
	val, ok := p[key]
	if !ok {
		return 0
	}
	res, _ := strconv.Atoi(val)
	return res
}

func (p Params) Bool(key string) bool {
	_, ok := p[key]
	return ok
}

func (p Params) Keys() (res []string) {
	for key, _ := range p {
		res = append(res, key)
	}
	return
}

type Request struct {
	Params Params
}
