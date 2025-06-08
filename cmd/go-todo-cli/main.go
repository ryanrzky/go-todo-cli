package main

var pathFileTodos string

func main() {
	todos := Todos{}
	storage := NewStorage[Todos](pathFileTodos)
	storage.Load(&todos)
	commandFlags := NewCommandFlags()
	commandFlags.Execute(&todos)
	storage.Save(todos)
}
