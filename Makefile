# change this
PATH_TODOS := "/Users/ryanrizky/todos.json"

APP_NAME := "todo"

build:
	go mod tidy
	go build -o ${APP_NAME} -ldflags "-X main.pathFileTodos=${PATH_TODOS}" ./cmd/go-todo-cli 
