/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:34
 */

package parser

import (
	"io"
	"io/ioutil"
	"todo/internal/context"
)

func Parser(reader io.Reader, param context.Params) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(data)
	}
	index := 0
	for len(data) > index {
		index = skip(index, data, skipSpace, skipSplit)
		if index >= len(data) {
			break
		}
		var key, val string
		key, val, index = getParam(data)
		param.Add(key, val)
	}
	return
}

func getParam(data []byte) (string, string, int) {
	var (
		key, val []byte
		i        int
	)
	for len(data) > i {
		v := data[i]
		if v == ' ' || v == '=' {
			break
		}
		key = append(key, v)
		i++
	}

	i = skip(i, data, skipEq)

	for len(data) > i {
		v := data[i]
		if v == '=' {
			break
		}
		val = append(val, v)
		i++
	}
	if len(val) == 0 {
		val = key
	}
	return string(key), string(val), i
}

func skip(i int, data []byte, skips ...skipFunc) int {
	for _, skip := range skips {
		for len(data) > i && skip(data[i]) {
			i++
		}
	}
	return i
}

type skipFunc func(byte) bool

func skipSpace(data byte) bool {
	return data == ' '
}

func skipSplit(data byte) bool {
	return data == '-'
}

func skipEq(data byte) bool {
	return data == '='
}
