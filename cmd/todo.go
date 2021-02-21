package main

import (
	"bytes"
	"fmt"
	"os"
	"todo/internal/app"
	"todo/internal/parser"
	"todo/internal/store"
)

func main() {
	var actionName string
	if len(os.Args) < 2 {
		actionName = "help"
	}
	actionName = os.Args[1]
	fs := store.FileStore{}
	fs.Init()
	e := app.NewEngine(parser.Parser, &fs)

	reader := bytes.NewBuffer(nil)
	for _, arg := range os.Args[2:] {
		reader.WriteString(arg)
	}
	buf := bytes.NewBuffer(nil)
	e.Run(actionName, reader, buf)
	fmt.Println(buf.String())
}
