/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:00
 */

package model

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

func (ts *Todos) Add(todo Todo) {
	*ts = append(*ts, todo)
}

type FilterFunc func(todo Todo) bool

func MergeFilter(filters ...FilterFunc) FilterFunc {
	return func(todo Todo) bool {
		for _, filter := range filters {
			if !filter(todo) {
				return false
			}
		}
		return true
	}
}

func (ts Todos) Filter(filter FilterFunc) (res Todos) {
	for _, t := range ts {
		if filter(t) {
			res = append(res, t)
		}
	}
	return
}

func (ts Todos) Size() int {
	return len(ts)
}
