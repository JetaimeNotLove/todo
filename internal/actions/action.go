package actions

import (
	"todo/internal/context"
)

var (
	actionMap map[string]Action
)

func init() {
	actionMap = make(map[string]Action)
}

func RegAction(name string, action Action) {
	actionMap[name] = action
}

func DoAction(name string, ctx *context.Context) error {
	action, ok := actionMap[name]
	if !ok {
		return ErrActionNotFound
	}
	return action(ctx)
}

type Action func(*context.Context) error
