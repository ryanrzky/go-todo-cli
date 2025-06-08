package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandFlags struct {
	Add    string
	Del    int
	List   bool
	Edit   string
	Toggle int
}

func NewCommandFlags() *CommandFlags {
	cf := CommandFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add new todo")
	flag.IntVar(&cf.Del, "del", -1, "Delete todo by index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.StringVar(&cf.Edit, "edit", "", "Edit todo by index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")

	flag.Parse()

	return &cf
}

func (cf *CommandFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.Print()

	case cf.Add != "":
		todos.Add(cf.Add)

	case cf.Del != -1:
		todos.Delete(cf.Del)

	case cf.Edit != "":
		splitParams := strings.SplitN(cf.Edit, ":", 2)
		if len(splitParams) != 2 {
			fmt.Println("Error: Invalid format to edit. Please use index:new_title")
		}

		index, err := strconv.Atoi(splitParams[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit")
			os.Exit(1)
		}

		todos.Edit(index, splitParams[1])
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)

	default:
		flag.PrintDefaults()
	}
}
