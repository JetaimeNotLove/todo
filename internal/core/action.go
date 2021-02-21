package core

var (
	actionMap map[string]Action
)

func init() {
	actionMap = make(map[string]Action)
}

func RegAction(name string, action Action) {
	actionMap[name] = action
}

func DoAction(name string, ctx *Context) {
	actionMap[name](ctx)
}

type Action func(*Context)
