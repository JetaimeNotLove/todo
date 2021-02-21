/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:00
 */

package core

type Status int

const (
	Done Status = 1
	Undo Status = 2
)

type Todo struct {
	Index   int
	Status  Status
	Content string
}

type Todos []Todo
