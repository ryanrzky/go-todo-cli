# Golang Todo CLI App

A simple command-line interface (CLI) application built with Go to manage your to-do list directly from the terminal.

## Features

* **Add Todos:** Quickly add new tasks to your list.
* **List Todos:** View all your pending and completed tasks.
* **Edit Todos:** Modify existing tasks.
* **Delete Todos:** Remove tasks you no longer need.
* **Toggle Status:** Mark tasks as complete or incomplete.

## Installation

To install the Golang Todo CLI App, follow these steps:

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/ryanrzky/go-todo-cli.git
    cd go-todo-cli
    ```

2.  **Edit the `Makefile`:**
    Open the `Makefile` in the project root. Locate the `PATH_TODOS` variable and set it to your desired file path to store todos data. This is where the data stored will be placed. For example:

    ```makefile
    # Makefile
    PATH_TODOS = "/Users/ryan"
    ```

3.  **Build:**
    Run the `make build` command from the project root. This will compile the application and move the executable to the `PATH_TODOS` directory you specified.

    ```bash
    make build
    ```

4. **Install**
    To enable the tools to be called from any location, move the binary file to `/usr/local/bin`.

    ```bash
    sudo mv ./todo /usr/local/bin/`
    ```

## Usage

Once installed, you can use the `todo` command (or whatever you named the executable in your `Makefile`) with the following arguments:

* **Add a new todo:**
    ```bash
    todo -add "Buy groceries"
    ```

* **List all todos:**
    ```bash
    todo -list
    ```

* **Edit a todo by index:**
    ```bash
    todo -edit "1: Call mom"
    ```
    (Here, `1` is the index of the todo you want to edit)

* **Delete a todo by index:**
    ```bash
    todo -del 2
    ```
    (Deletes the todo at index `2`)

* **Toggle a todo's completion status by index:**
    ```bash
    todo -toggle 3
    ```
    (Toggles the status of the todo at index `3`)