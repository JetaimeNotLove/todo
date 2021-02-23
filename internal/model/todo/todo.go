/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:00
 */

package todo

type Status int

func (s Status) String() string {
	switch s {
	default:
		return ""
	case Done:
		return "Done"
	case Undo:
		return "Undo"
	}
}

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

func (ts Todos) Get(index int) *Todo {
	for i, t := range ts {
		if t.Index == index {
			return &ts[i]
		}
	}
	return nil
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
