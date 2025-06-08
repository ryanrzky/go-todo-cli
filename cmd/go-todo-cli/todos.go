package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Complete    bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:       title,
		Complete:    false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) ValidateIndex(index int) error {
	if index < 0 || index > len(*todos) {
		err := errors.New("Error: Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos

	err := t.ValidateIndex(index)
	if err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos

	err := t.ValidateIndex(index)
	if err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) Toggle(index int) error {
	t := *todos

	err := t.ValidateIndex(index)
	if err != nil {
		return err
	}

	isComplete := t[index].Complete

	if !isComplete {
		t[index].Complete = true
		completeTime := time.Now()
		t[index].CompletedAt = &completeTime
	} else {
		t[index].Complete = false
		var completeTime *time.Time = nil
		t[index].CompletedAt = completeTime
	}

	return nil
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Id", "Title", "Complete", "Created At", "Completed At")

	for index, t := range *todos {
		complete := "❌"
		completedAt := ""

		if t.Complete {
			complete = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, complete, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
