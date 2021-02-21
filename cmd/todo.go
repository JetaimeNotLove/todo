package main

import (
	"bytes"
	"os"
	"todo/internal/core"
	"todo/internal/parser"
	"todo/internal/store"
)

func main() {
	var actionName string
	if len(os.Args) < 2 {
		actionName = "help"
	}
	actionName = os.Args[1]
	e := core.NewEngine(parser.Parser, store.FileStore{})

	reader := bytes.NewBuffer(nil)
	for _, arg := range os.Args[2:] {
		reader.WriteString(arg)
	}
	write := bytes.NewBuffer(nil)
	e.Run(actionName, reader, write)
}
