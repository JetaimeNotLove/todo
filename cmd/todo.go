package main

import (
	"bytes"
	"fmt"
	"os"
	"todo/internal/app"
	"todo/internal/dao/filedao"
	"todo/internal/parser"
)

func main() {
	var actionName string
	if len(os.Args) < 2 {
		actionName = "help"
	}
	actionName = os.Args[1]
	fs := filedao.FileStore{}
	fs.Init()
	e := app.NewApp(parser.Parser, &fs)

	reader := bytes.NewBuffer(nil)
	for _, arg := range os.Args[2:] {
		reader.WriteString(arg)
	}
	buf := bytes.NewBuffer(nil)
	e.Run(actionName, reader, buf)
	fmt.Println(buf.String())
}
