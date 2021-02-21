/**
 * @Author: Huyantian
 * @Date: 2021/2/21 下午3:00
 */

package core

type QueryParam struct {
	Status Status
}

type QueryFunc func(param *QueryParam)

type UpdateParam struct {
	M map[string]interface{}
}

type UpdateFunc func(param *UpdateParam)

type TodoStore interface {
	List(ctx Context, queries ...QueryFunc) (Todos, error)
	Create(ctx Context, todo Todo) error
	Update(ctx Context, index int, updates ...UpdateFunc) error
}
