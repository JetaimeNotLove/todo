package main

import (
	"bytes"
	"fmt"
	"os"
	"todo/internal/app"
	"todo/internal/dao"
	"todo/internal/dao/filedao"
	"todo/internal/parser"
)

func main() {

	// 初始化存储
	filedao.Init()
	dao.SetTodo(&filedao.TodoDao{})
	dao.SetUser(&filedao.UserDao{})

	e := app.NewApp(parser.Parser)

	input := bytes.NewBuffer(nil)
	GetInput(input)
	output := bytes.NewBuffer(nil)
	e.Run(GetAction(), input, output)
	fmt.Println(output.String())
}

func GetInput(reader *bytes.Buffer) {
	for _, arg := range os.Args[2:] {
		reader.WriteString(arg)
	}
}

func GetAction() string {
	var actionName string
	if len(os.Args) < 2 {
		actionName = "help"
	}
	actionName = os.Args[1]
	return actionName
}
